# CCV Distribution
                                                                             
```                                                                          
              ┌─for each block────────────────────────────────────┐     
          ┌─  │┌─────────────────┐             ┌───────────────┐  │ -- Params -- 
          │   ││consumer-chain   │             │fee-collector  │  │  [P1] BlocksPerDistributionTransmission                  
          │   ││mints & collects ├──sends-to──▶│holding address│  │                                                         
          │   ││reward tokens    │             └───────┬───────┘  │ -- Keys --                                        
          │   │└─────────────────┘                     │          │  [K1] LastDistributionTransmission
          │   │                                        │          │  [K2] DistributionValidatorHoldingPool
          │   │                                        ▼          └──────────┐
  Consumer│   │                        ┌──AllocateTokens───────────────────┐ │
  Chain  ─┤   │ tendermint votes,      │distribute funds from fee-collector│ │
          │   │ power, & proposer ────▶│to per-validator holding pool      │ │
          │   │                        │using AllocateTokens-              │ │
          │   │                        │ToProviderValidatorHoldingPool.    │ │
          │   │                        │(pools held with key prefix: [K2]) │ │                          
          │   │                        └───────────────────────────────────┘ │
          │   └─────────────────┬────────────────────────────────────────────┘
          │                     │                                  ▲ 
          │              Wait [P1] Blocks                          │ 
          │                     │                                  │               
          │                     ▼                                  │               
          │   ┌──────────────────────────────────────────┐         │            
          │   │Combine all rewards held in all per-      │         │            
          │   │validator holding pools into a single pool│  send back remainder
          │   │(ProviderPoolTokens) while recording      │  to fee-collector for
          │   │the fraction of this pool owed to each    │  use in next block
          │   │validate into a ProviderPoolWeights object│         │            
          │   └───────────┬───────────┬──────────────────┘         │            
          │               │           │    ┌───────────────────────┴─────┐    
          │     ┌─────────┴─────────┐ └────┤Truncate ProviderPoolTokens  │ 
          │     │ProviderPoolWeights│      │converting: DecCoins -> Coins│
          │     └─────────┬─────────┘      └─────────────┬───────────────┘  
          │               │                              │                
          └─              ▼                              ▼  
          ┌─    transmit to provider            transmit to provider                            
  Relayer─┤     chain on CCV ordered            chain on unordered                                      
          └─    IBC channel                     IBC channel (ICS-20)                                     
          ┌─             │                               │
          │              │                    ┌──────────┘
          │              │                    │      
          │     Wait for both packets to be received
          │              │                          
  Provider│    ┌─initialize────────────────────────────────────────────┐
  Chain  ─┤    │QualifiedTotalWeight := ProviderPoolWeights.TotalWeight│
          │    │DisqualiedPool       := 0                              │
          │    └─────────┬─────────────────────────────────────────────┘
          │              │
          │              ▼
          │    ┌─for each validator[i] in ProviderPool───────────────────────────────┐
          │    │┌──────────────┐    ┌──────────────────────────────────────────────┐ │
          │    ││Does validator│    │Validator forfeits rewards:                   │ │
          │    ││still exist?  │    │                                              │ │
          │    │└──┬───┬───────┘    │DisqualifiedPool = ProviderPoolTokens         │ │     
          │    │   │   │            │                      * ProviderPoolWeights[i]│ │     
          │    │   │   └──yes──────▶│QualifiedTotalWeight -= ProviderPoolWeights[i]│ │              
          │    │   │                └──────────────────────────────────────────────┘ │
          │    │   │                ┌───────────────────────┐                        │
          │    │   └───────no──────▶│added to array         │                        │
          │    │                    │of qualified validators│                        │
          │    │                    └───────────────────────┘                        │
          │    └─────────┬───────────────────────────────────────────────────────────┘
          │              ├───▶if no qualified validators send DisqualifiedPool
          │              │    to provider community pool (edge case)       
          │              ▼                                                                         
          │    ┌─for each qualified validator[i]─────────────────────────────────────┐ 
          │    │┌──────────────────────────────────────┐┌───────────────────────────┐│ 
          │    ││Calculate rewards:                    ││Final distribution using   ││ 
          │    ││TW := ProviderPoolWeights.TotalWeight ││AllocateTokensToValidator: ││ 
          │    ││W  := ProviderPoolWeights[i]          ││ -> delegator rewards      ││ 
          │    ││ValRewards := ProviderPoolTokens*W/TW ││ -> validator commission   ││ 
          │    ││              + DisqualifiedPool      ││                           ││ 
          │    ││              * W/QualifiedTotalWeight││                           ││   
          │    │└──────────────────────────────────────┘└───────────────────────────┘│ 
          └─   └─────────────────────────────────────────────────────────────────────┘ 
                                                                                               
```        
           
           
           
           
           
           
           
           