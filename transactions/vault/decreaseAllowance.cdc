import FiatToken from 0x{{.FiatToken}}
import FiatTokenInterface from 0x{{.FiatTokenInterface}}

transaction(toResourceId: UInt64, delta: UFix64) {

    prepare (fromAcct: AuthAccount) {
        // Get a reference to the signer's stored vault
        let vaultRef = fromAcct.borrow<&FiatToken.Vault>(from: FiatToken.VaultStoragePath)
            ?? panic("Could not borrow reference to the owner's Vault!")

        vaultRef.decreaseAllowance(resourceId: toResourceId, decrement: delta)
    }
}