# ASSIGNMENT
STEPS FOR INSTALLING AND INSTATIATING CHAINCODE ON HYPERLEDGERFABRIC 2.2
==============================================================================

// CREATE NFS SERVER

//START UP FABRIC CA SERVER

//GENERATE CERTIFICATE FOR PEER AND ORDERER FOR EVERY ORGANISATION

// GENERATE GENESIS BLOCK AND CHANNEL TRANSACTION

configtxgen -profile TwoOrgsOrdererGenesis -channelID system-channel -outputBlock ./system-genesis-block/genesis.block

// CREATE CHANNEL

configtxgen -profile TwoOrgsChannel -outputCreateChannelTx ./channel-artifacts/${CHANNEL_NAME}.tx -channelID $CHANNEL_NAME

// JOIN CHANNEL FOR ALL PEER
 
//PACKAGE THE CHAINCODE

peer lifecycle chaincode package ${CC_PATH}/${CC_NAME}.tar.gz --path ${CC_PATH}/$CC_NAME --label ${CC_NAME}

// TRANSFER THE FILE IN ALL WORKER NODE

// NOW INSTALL THE CHAINCODE

peer lifecycle chaincode install ${CC_PATH}/$CC_NAME.tar.gz // INSTALL THE CHAINCODE ON ALL NODE

//YOU WILL GET PACKAGE IDENTIFIER AFTER INSTALLATION

// NOW APPROVE THE CHAINCODE

peer lifecycle chaincode approveformyorg -o orderer.example.com:___ --tls --cafile $ORDERER_CA --channelID $CHANNEL_NAME --name $CC_NAME --version 1 --package-id $PACKAGE_ID --sequence 1 --init-required --signature-policy "OR ('Org1MSP.peer','Org2MSP.peer','Org3MSP.peer')" 

// APPROVE FOR ALL ORGANISATION
a
// NOW COMMIT THE CHANCODE

peer lifecycle chaincode commit -o orderer.example.com:___ --tls --cafile $ORDERER_CA --channelID $CHANNEL_NAME --name $CC_NAME $PEER_CONN_PARMS --version 1 --sequence 1 --init-required --signature-policy "OR ('Org1MSP.peer','Org2MSP.peer','Org3MSP.peer')" 

// INVOKE CHAINCODE
peer chaincode invoke -o orderer.example.com:____ --tls --cafile $ORDERER_CA -C mychannel -n ${CC_NAME} $PEER_CONN_PARMS -c ${fcn_call}

-----------------------------------------------------------------------------------------------------------------------------------------

CRYPTOGEN AND CONFIGTXGEN

Cryptogen is a tool used to generate cryptographic material for use with Hyperledger Fabric. This material includes certificates, keys, and other cryptographic artifacts. Cryptogen takes as input a configuration file that specifies the organizations and entities that will be participating in the Hyperledger Fabric network. The output of Cryptogen is a set of directories that contain the generated cryptographic material
------------------------------------------YAML--------------------------------------------------
Organizations:
- Name: Org1MSP
  ID: Org1MSP
  MSPDir: crypto-config/peerOrganizations/org1.test.com/msp
  adminPrivateKeyPEM: crypto-config/peerOrganizations/org1.test.com/admin/key.pem
  adminCertificatePEM: crypto-config/peerOrganizations/org1.test.com/admin/cert.pem

  Users:
  - Name: User1@org1.test.com
    KeyFile: crypto-config/peerOrganizations/org1.test.com/users/User1@org1.test.com/key.pem
    CertificatePEM: crypto-config/peerOrganizations/org1.test.com/users/User1@org1.test.com/cert.pem

-------------------------------------------------------------------------------------------------------

Configtxgen is a tool used to generate channel configuration artifacts for use with Hyperledger Fabric. Channel configuration artifacts are used to define the configuration of a channel, including the organizations that are participating in the channel, the policies for accessing the channel, and the orderers that are responsible for ordering transactions on the channel. Configtxgen takes as input a configuration file that specifies the channel configuration. The output of Configtxgen is a set of files that contain the generated channel configuration artifacts.

---------------------------------------YAML----------------------------------------------
Organizations:
- Name: Org1MSP
  ID: Org1MSP
  MSPDir: crypto-config/peerOrganizations/org1.test.com/msp

  AnchorPeers:
  - Host: peer0.org1.test.com
    Port: 7051

    Peers:
    - Host: peer1.org1.test.com
      Port: 7051
-----------------------------------------------------------------------------------------------


