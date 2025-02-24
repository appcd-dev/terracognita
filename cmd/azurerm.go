package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/cycloidio/terracognita/azurerm"
	"github.com/cycloidio/terracognita/log"
)

var (
	azurermCmd = &cobra.Command{
		Use:   "azurerm",
		Short: "Terracognita reads from Azure and generates hcl resources and/or terraform state",
		Long:  "Terracognita reads from Azure and generates hcl resources and/or terraform state",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			err := preRunEOutput(cmd, args)
			if err != nil {
				return err
			}
			viper.BindPFlag("client-id", cmd.Flags().Lookup("client-id"))
			viper.BindPFlag("client-secret", cmd.Flags().Lookup("client-secret"))
			viper.BindPFlag("environment", cmd.Flags().Lookup("environment"))
			viper.BindPFlag("resource-group-name", cmd.Flags().Lookup("resource-group-name"))
			viper.BindPFlag("subscription-id", cmd.Flags().Lookup("subscription-id"))
			viper.BindPFlag("tenant-id", cmd.Flags().Lookup("tenant-id"))
			viper.BindPFlag("tags", cmd.Flags().Lookup("tags"))

			return nil
		},
		PostRunE: postRunEOutput,
		RunE: func(cmd *cobra.Command, args []string) error {
			logger := log.Get()
			logger = logger.With("func", "cmd.azure.RunE")
			// Validate required flags
			if err := requiredStringFlags(
				"client-id", "client-secret", "subscription-id", "tenant-id",
			); err != nil {
				return err
			}
			if len(viper.GetStringSlice("resource-group-name")) == 0 {
				return fmt.Errorf("the flag 'resource-group-name' is required")
			}

			azureRMP, err := azurerm.NewProvider(
				cmd.Context(),
				viper.GetString("client-id"),
				viper.GetString("client-secret"),
				viper.GetString("environment"),
				viper.GetStringSlice("resource-group-name"),
				viper.GetString("subscription-id"),
				viper.GetString("tenant-id"),
			)
			if err != nil {
				return err
			}

			tags, err := initializeTags("tags")
			if err != nil {
				return err
			}

			err = importProvider(cmd.Context(), logger, azureRMP, tags)
			if err != nil {
				return err
			}

			return nil
		},
	}
)

func init() {
	azurermCmd.AddCommand(azurermResourcesCmd)

	// Required flags
	azurermCmd.Flags().String("client-id", "", "Client ID (required)")
	azurermCmd.Flags().String("client-secret", "", "Client Secret (required)")
	azurermCmd.Flags().StringSlice("resource-group-name", nil, "Resource Group Names (required)")
	azurermCmd.Flags().String("subscription-id", "", "Subscription ID (required)")
	azurermCmd.Flags().String("tenant-id", "", "Tenant ID (required)")

	azurermCmd.Flags().StringSliceVarP(&tags, "tags", "t", []string{}, "List of tags to filter with format 'NAME:VALUE'")

	// Optional flags
	azurermCmd.Flags().String("environment", "public", "Environment")
}
