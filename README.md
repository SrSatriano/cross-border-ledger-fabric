# Ledger Privado para Rastreabilidade Cross-Border

Blockchain permissionada Hyperledger Fabric para faturas e fluxos internacionais auditáveis.

## Stack

- Hyperledger Fabric, Go chaincode, Docker Swarm

## Arquitetura de canais

```
Orderer
   │
├── Channel BR-Ops (Org A, Org B)
├── Channel EU-Tax (Org B, Org C)
└── Channel Audit (todas orgs read-only)
```

## Políticas de endosso

- Transação exige `AND(OrgA.peer, OrgB.peer)`
- Leitura pública restrita ao canal

## Bootstrap B2B

```bash
./network/scripts/bootstrap.sh
docker stack deploy -c docker-compose.swarm.yml fabric
```

Documentação: [docs/ENDORSEMENT.md](docs/ENDORSEMENT.md) | [docs/CHANNELS.md](docs/CHANNELS.md)

## Chaincode

`chaincode/invoice/` — registrar fatura, reter imposto, hash imutável.
