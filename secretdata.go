/* Copyright © 2020. Financelime, https://financelime.com. All rights reserved.
   Author: DmAlix. Contacts: <dmalix@financelime.com>, <dmalix@yahoo.com>
   License: GNU General Public License v3.0, https://www.gnu.org/licenses/gpl-3.0.html */

package secretdata

type Cipher struct {
	SecretKey string
}

func NewSecretData(
	SecretKey string) *Cipher {
	return &Cipher{
		SecretKey: SecretKey,
	}
}
