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
	"github.com/spf13/cobra"
)

var syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "Update bespoke from GitHub",
	Run: func(cmd *cobra.Command, args []string) {
		execSync()
	},
}

func execSync() {
	// TODO: download hooks asset (latest or user selected), and extract to $XDG_CONFIG_HOME/bespoke/hooks
}

func init() {
	rootCmd.AddCommand(syncCmd)
}
