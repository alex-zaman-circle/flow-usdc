// Masterminter uses this to configure which minter the minter controller manages

import FiatToken from 0x{{.FiatToken}}
import FiatTokenInterface from 0x{{.FiatTokenInterface}}
import OnChainMultiSig from 0x{{.OnChainMultiSig}}

transaction (keyListIndex: Int, sig: [UInt8], addr: Address, method: String, minter: UInt64, minterController: UInt64) {
    prepare(oneOfMultiSig: AuthAccount) {
    }

    execute {
        // Get the recipient's public account object
        let masterMinterOwnerAcct = getAccount(addr)

        // Get a allowance reference to the fromAcct's vault 
        let pubSigRef = masterMinterOwnerAcct.getCapability(FiatToken.MasterMinterPubSigner)
            .borrow<&FiatToken.MasterMinter{OnChainMultiSig.PublicSigner}>()
            ?? panic("Could not borrow master minter pub sig reference")
            
        let minterArg = OnChainMultiSig.PayloadArg(t: Type<UInt64>(), v: minter);
        let minterControllerArg = OnChainMultiSig.PayloadArg(t: Type<UInt64>(), v: minterController);
        let p = OnChainMultiSig.PayloadDetails(method: method, args: [minterArg, minterControllerArg])

        // pub fun addNewPayload(payload: OnChainMultiSig.PayloadDetails, keyListIndex: Int, payloadSig: [UInt8]): UInt64 {
        return pubSigRef.addNewPayload(payload: p, keyListIndex: keyListIndex, sig: sig) 
    }
}
