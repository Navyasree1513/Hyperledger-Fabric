const express = require("express");
const { Gateway, Wallets } = require("fabric-network");
const path = require("path");
const fs = require("fs");

const app = express();
app.use(express.json());

const ccpPath = path.resolve(__dirname, "..", "fabric-samples", "test-network", "organizations", "peerOrganizations", "org1.example.com", "connection-org1.json");

async function getContract() {
    const ccp = JSON.parse(fs.readFileSync(ccpPath, "utf8"));
    const wallet = await Wallets.newFileSystemWallet("./wallet");

    const gateway = new Gateway();
    await gateway.connect(ccp, {
        wallet,
        identity: "appUser",
        discovery: { enabled: true, asLocalhost: true },
    });

    const network = await gateway.getNetwork("mychannel");
    return network.getContract("assetcc");
}

app.post("/create", async (req, res) => {
    try {
        const contract = await getContract();
        const { dealerId, msisdn, mpin, balance, status, transAmt, transType, remarks } = req.body;
        await contract.submitTransaction("CreateAsset", dealerId, msisdn, mpin, balance, status, transAmt, transType, remarks);
        res.json({ status: "success" });
    } catch (err) {
        res.status(500).send(err.message);
    }
});

app.get("/read/:id", async (req, res) => {
    try {
        const contract = await getContract();
        const result = await contract.evaluateTransaction("ReadAsset", req.params.id);
        res.json(JSON.parse(result.toString()));
    } catch (err) {
        res.status(500).send(err.message);
    }
});

app.listen(3000, () => console.log("REST API running on port 3000"));
