# `/x/rps` - Rock, Paper or Scissors Module

This directory contains the code for your chain custom modules.


## State objects

### Game

- gameNumber (uint)
- playerA (adress)
- playerB (address)
- status (string)
- rounds (uint)
- playerAMoves ([]string)
- playerBMoves ([]string)
- score ([2]uint)

### Param


## Msg Service

### MsgCreateGame

- creator (address - signer)
- oponent (address)
- rounds (uint)

### MsgMakeMove
- player (address - signer)
- gameNumber (uint)
- move (string Rock | Paper | Scissors) 

## Query Service

### GetGame

- gameNumber (uint)

### GetParams

## Events

### EventCreateGame

- gameNumber (uint)
- playerA (address)
- playerB (address)

### EventEndGame

- gameNumber (uint)
- status (string)

### EventMakeMove

- gameNumber (uint)
- player (address)
- move (string Rock | Paper | Scissors)

