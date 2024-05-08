/* Copyright © 2024
 *      Delusoire <deluso7re@outlook.com>
 *
 * This file is part of bespoke/cli.
 *
 * bespoke/cli is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * bespoke/cli is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with bespoke/cli. If not, see <https://www.gnu.org/licenses/>.
 */

package cmd

import (
	"bespoke/module"
	"log"

	"github.com/spf13/cobra"
)

var pkgCmd = &cobra.Command{
	Use:   "pkg [action]",
	Short: "Manage modules",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var pkgInstallCmd = &cobra.Command{
	Use:   "install [murl]",
	Short: "Install module",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		metadataURL := args[0]
		if err := module.InstallModuleMURL(metadataURL); err != nil {
			log.Fatalln(err.Error())
		}
	},
}

var pkgDeleteCmd = &cobra.Command{
	Use:   "rem [id]",
	Short: "Uninstall module",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		identifier := module.NewStoreIdentifier(args[0])
		if err := module.DeleteModule(identifier); err != nil {
			log.Fatalln(err.Error())
		}
	},
}

var pkgEnableCmd = &cobra.Command{
	Use:   "enable [id]",
	Short: "Enable installed module",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		identifier := module.NewStoreIdentifier(args[0])
		if err := module.ToggleModuleInVault(identifier, true); err != nil {
			log.Fatalln(err.Error())
		}
	},
}

var pkgDisableCmd = &cobra.Command{
	Use:   "disable [id]",
	Short: "Disable installed module",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		identifier := module.NewStoreIdentifier(args[0])
		if err := module.ToggleModuleInVault(identifier, false); err != nil {
			log.Fatalln(err.Error())
		}
	},
}

func init() {
	rootCmd.AddCommand(pkgCmd)

	pkgCmd.AddCommand(pkgInstallCmd, pkgDeleteCmd, pkgEnableCmd, pkgDisableCmd)
}
