package cmd

import (
	"fmt"
	"strings"

	"github.com/alexellis/bazaar/cmd/apps"
	"github.com/spf13/cobra"
)

func MakeInstall() *cobra.Command {
	var command = &cobra.Command{
		Use:   "install",
		Short: "Install Kubernetes apps from helm charts or YAML files",
		Long: `Install Kubernetes apps from helm charts or YAML files using the "install" 
command. Helm 2 is used by default unless a --helm3 flag exists and is passed. 
You can also find the post-install message for each app with the "info" 
command.`,
		Example: `  bazaar install
  bazaar install openfaas --helm3 --gateways=2
  bazaar app info inlets-operator`,
		SilenceUsage: false,
	}

	var install = &cobra.Command{
		Use:   "install",
		Short: "Install a Kubernetes app",
		Example: `  bazaar install [APP]
  bazaar install openfaas --help
  bazaar install inlets-operator --token-file $HOME/do
  bazaar install --help`,
		SilenceUsage: true,
	}

	install.PersistentFlags().String("kubeconfig", "kubeconfig", "Local path for your kubeconfig file")

	install.RunE = func(command *cobra.Command, args []string) error {

		if len(args) == 0 {
			fmt.Printf("You can install: %s\n%s\n\n", strings.TrimRight("\n - "+strings.Join(getApps(), "\n - "), "\n - "),
				`Run bazaar install NAME --help to see configuration options.`)
			return nil
		}

		return nil
	}

	command.AddCommand(install)
	install.AddCommand(apps.MakeInstallOpenFaaS())
	install.AddCommand(apps.MakeInstallMetricsServer())
	install.AddCommand(apps.MakeInstallInletsOperator())
	install.AddCommand(apps.MakeInstallCertManager())
	install.AddCommand(apps.MakeInstallOpenFaaSIngress())
	install.AddCommand(apps.MakeInstallNginx())
	install.AddCommand(apps.MakeInstallChart())
	install.AddCommand(apps.MakeInstallLinkerd())
	install.AddCommand(apps.MakeInstallCronConnector())
	install.AddCommand(apps.MakeInstallKafkaConnector())
	install.AddCommand(apps.MakeInstallMinio())
	install.AddCommand(apps.MakeInstallPostgresql())
	install.AddCommand(apps.MakeInstallKubernetesDashboard())
	install.AddCommand(apps.MakeInstallIstio())
	install.AddCommand(apps.MakeInstallCrossplane())
	install.AddCommand(apps.MakeInstallMongoDB())
	install.AddCommand(apps.MakeInstallRegistry())
	install.AddCommand(apps.MakeInstallRegistryIngress())

	command.AddCommand(MakeInfo())

	return command
}

func getApps() []string {
	return []string{"openfaas",
		"nginx-ingress",
		"cert-manager",
		"openfaas-ingress",
		"inlets-operator",
		"metrics-server",
		"chart",
		"tiller",
		"linkerd",
		"cron-connector",
		"kafka-connector",
		"minio",
		"postgresql",
		"kubernetes-dashboard",
		"istio",
		"crosspane",
		"mongodb",
		"docker-registry",
		"docker-registry-ingress",
	}
}
