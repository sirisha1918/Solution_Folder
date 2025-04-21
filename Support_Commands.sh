# To set the environment for Org1: 
export FABRIC_CFG_PATH=${PWD}/configtx  
source ./scripts/setOrgPeerContext.sh 1 

peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls --cafile $ORDERER_CA -C mychannel -n paymentscc --peerAddresses localhost:7051 --tlsRootCertFiles $PEER0_ORG1_CA --peerAddresses localhost:9051 --tlsRootCertFiles $PEER0_ORG2_CA -c '{"Args":["RegisterAccount", "account1", "Alice", "1000"]}' 

peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls --cafile $ORDERER_CA -C mychannel -n paymentscc --peerAddresses localhost:7051 --tlsRootCertFiles $PEER0_ORG1_CA --peerAddresses localhost:9051 --tlsRootCertFiles $PEER0_ORG2_CA -c '{"Args":["RegisterAccount", "account2", "Bob", "500"]}' 
 
peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls --cafile $ORDERER_CA -C mychannel -n paymentscc --peerAddresses localhost:7051 --tlsRootCertFiles $PEER0_ORG1_CA --peerAddresses localhost:9051 --tlsRootCertFiles $PEER0_ORG2_CA -c '{"Args":["TransferFunds", "account1", "account2", "200"]}' 
 
peer chaincode query -C mychannel -n paymentscc \-c '{"Args":["GetBalance", "account1"]}' 

peer chaincode query -C mychannel -n paymentscc \ -c '{"Args":["GetBalance", "account2"]}' 

 
