
# Open Plasma Architecture

  

## Description

Open Plasma is composed of multiple applications. The core part of them:

  - plasma operator
  - plasma block explorer
  - plasma smart contract

We use [More Viable Plasma](https://ethresear.ch/t/more-viable-plasma/2160) implementation, then plasma operator builds plasma chain, deliver blocks to the client application (block explorer) and hashes of blocks to plasma contract.

Block explorer watches out for the operator. If the operator cheats or does something wrong, plasma contract guarantees a refund of assets from plasma contract to mainnet.

Plasma contract provides following operations: deposits, exits (including special exits for in-flight transactions), keeps enumerated plasma block hashes. We do not challenge each action of the operator on plasma contract, because the operator has "silver bullet" against any challenges: block withholding. There is no way to prove blocks unavailability. So, fair clients have an an opportunity to withdraw their assets from plasma to mainnet.

## Plasma Operator

  

### Block structure

  
  

## Plasma client application

  

## Ethereum Smart Contract

  

## Appendixes

  

### Exit game

  

We suppose to use standard [More Viable Plasma](https://ethresear.ch/t/more-viable-plasma/2160) exit game with some additions. As we are going to support atomic swaps with multiple transaction source owners, one of a participant may have not all signatures of the transaction. So, we need to set up one more exit game branch with signatures collection.

  

![exit game schema](https://raw.githubusercontent.com/BANKEX/plasma-research/master/docs/assets/plasma_exit_game.svg?sanitize=true)
