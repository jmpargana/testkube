package pro

import (
	"strings"

	"github.com/pterm/pterm"

	"github.com/spf13/cobra"

	"github.com/kubeshop/testkube/cmd/kubectl-testkube/commands/common"
	"github.com/kubeshop/testkube/cmd/kubectl-testkube/config"
	"github.com/kubeshop/testkube/pkg/ui"
)

func NewDisconnectCmd() *cobra.Command {

	var opts common.HelmOptions

	cmd := &cobra.Command{
		Use:     "disconnect",
		Aliases: []string{"d"},
		Short:   "Switch back to Testkube OSS mode, based on active .kube/config file",
		Run: func(cmd *cobra.Command, args []string) {

			ui.H1("Disconnecting your Pro environment:")
			ui.Paragraph("Rolling back to your clusters testkube OSS installation")
			ui.Paragraph("If you need more details click into following link: " + docsUrl)
			ui.H2("You can safely switch between connecting Pro and disconnecting without losing your data.")

			cfg, err := config.Load()
			if err != nil {
				pterm.Error.Printfln("Failed to load config file: %s", err.Error())
				return
			}

			client, _, err := common.GetClient(cmd)
			ui.ExitOnError("getting client", err)

			info, err := client.GetServerInfo()
			firstInstall := err != nil && strings.Contains(err.Error(), "not found")
			if err != nil && !firstInstall {
				ui.Failf("Can't get testkube cluster information: %s", err.Error())
			}
			var apiContext string
			if actx, ok := contextDescription[info.Context]; ok {
				apiContext = actx
			}
			var clusterContext string
			if cfg.ContextType == config.ContextTypeKubeconfig {
				clusterContext, err = common.GetCurrentKubernetesContext()
				if err != nil {
					pterm.Error.Printfln("Failed to get current kubernetes context: %s", err.Error())
					return
				}
			}

			// TODO: implement context info
			ui.H1("Current status of your Testkube instance")

			summary := [][]string{
				{"Testkube mode"},
				{"Context", apiContext},
				{"Kubectl context", clusterContext},
				{"Namespace", cfg.Namespace},
				{ui.Separator, ""},

				{"Testkube is connected to Pro organizations environment"},
				{"Organization Id", info.OrgId},
				{"Environment Id", info.EnvId},
			}

			ui.Properties(summary)

			if ok := ui.Confirm("Shall we disconnect your Pro environment now?"); !ok {
				return
			}

			ui.NL(2)

			spinner := ui.NewSpinner("Disonnecting from Testkube Pro")

			err = common.HelmUpgradeOrInstalTestkube(opts)
			ui.ExitOnError("Installing Testkube Pro", err)
			spinner.Success()

			// let's scale down deployment of mongo
			if opts.MongoReplicas > 0 {
				spinner = ui.NewSpinner("Scaling up MongoDB")
				common.KubectlScaleDeployment(opts.Namespace, "testkube-mongodb", opts.MongoReplicas)
				spinner.Success()
			}
			if opts.MinioReplicas > 0 {
				spinner = ui.NewSpinner("Scaling up MinIO")
				common.KubectlScaleDeployment(opts.Namespace, "testkube-minio-testkube", opts.MinioReplicas)
				spinner.Success()
			}
			if opts.DashboardReplicas > 0 {
				spinner = ui.NewSpinner("Scaling up Dashbaord")
				common.KubectlScaleDeployment(opts.Namespace, "testkube-dashboard", opts.DashboardReplicas)
				spinner.Success()
			}

			ui.NL()
			ui.Success("Disconnect finished successfully")
			ui.NL()
			ui.ShellCommand("You can now open your local Dashboard and validate the successfull disconnect", "testkube dashboard")
		},
	}

	// populate options
	common.PopulateHelmFlags(cmd, &opts)
	cmd.Flags().IntVar(&opts.MinioReplicas, "minio-replicas", 1, "MinIO replicas")
	cmd.Flags().IntVar(&opts.MongoReplicas, "mongo-replicas", 1, "MongoDB replicas")
	cmd.Flags().IntVar(&opts.DashboardReplicas, "dashboard-replicas", 1, "Dashboard replicas")
	return cmd
}
