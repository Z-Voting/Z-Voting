# Create Unique Id
```bash
peer chaincode query -n zvoting -c '{"Args":["generateUID"]}' -C myc
```

# Create Election
```bash
peer chaincode invoke -n zvoting -c '{"Args":["createElection", "cse election", "3600"]}' -C myc
```

# Get Elections
```bash
peer chaincode query -n zvoting -c '{"Args":["getElections"]}' -C myc
```

# Add Candidate
```bash
peer chaincode invoke -n zvoting -c '{"Args":["addCandidate", "Romana Mahjabin Eshita", "Rose", "rose.jpg", "LnfgDsc2WD8F2qNfHK5a"]}' -C myc
```

```bash
peer chaincode invoke -n zvoting -c '{"Args":["addCandidate", "Anwer H Anik", "Horlicks", "horlicks.jpg", "LnfgDsc2WD8F2qNfHK5a"]}' -C myc
```

# Register Voter
```bash
peer chaincode invoke -n zvoting -c '{"Args":["registerVoter", "Tanmoy Krishna Das", "tkd@gmail.com", "1", "1", "1", "LnfgDsc2WD8F2qNfHK5a"]}' -C myc
```

```bash
peer chaincode invoke -n zvoting -c '{"Args":["registerVoter", "Anam Ibna Harun", "anam@gmail.com", "1", "1", "1", "LnfgDsc2WD8F2qNfHK5a"]}' -C myc
```

# Get Candidates
```bash
peer chaincode query -n zvoting -c '{"Args":["getCandidates", "LnfgDsc2WD8F2qNfHK5a"]}' -C myc
```

# Get Login Challenge
```bash
peer chaincode query -n zvoting -c '{"Args":["getLoginChallenge"]}' -C myc
```

# Voter Login
```bash
peer chaincode query -n zvoting -c '{"Args":["voterLogin", "tkd@gmail.com", "4", "22","22","22", "1","1","1","2"]}' -C myc
```

```bash
peer chaincode query -n zvoting -c '{"Args":["voterLogin", "anam@gmail.com", "4", "22","22","22", "1","1","1","2"]}' -C myc
```


# Start Election
```bash
peer chaincode invoke -n zvoting -c '{"Args":["startElection", "LnfgDsc2WD8F2qNfHK5a"]}' -C myc
```


# Cast vote
```bash
peer chaincode invoke -n zvoting -c '{"Args":["castVote", "VUWSP2NcHciWvqZTa2N9", "[1,0]"]}' -C myc
```

```bash
peer chaincode invoke -n zvoting -c '{"Args":["castVote", "5RxRTZHWUsaD6HEdz0Th", "[1,0]"]}' -C myc
```

# Calculate Result
```bash
peer chaincode query -n zvoting -c '{"Args":["calculateResult", "LnfgDsc2WD8F2qNfHK5a"]}' -C myc
```

# Delete Everything

```bash
peer chaincode invoke -n zvoting -c '{"Args":["deleteAll"]}' -C myc
```

