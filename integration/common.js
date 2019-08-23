const DEVNET_ADDRESS = '0x9CA9d2D5E04012C9Ed24C0e513C9bfAa4A2dD77f'
const port = process.env.ETH_HTTP_PORT || `18545`

const abort = message => error => {
  console.error(message)
  console.error(error)
  process.exit(1)
}

// wrapper for main truffle script functions
const scriptRunner = main => async callback => {
  try {
    await main()
    callback()
  } catch (error) {
    callback(error)
  }
}

module.exports = { abort, DEVNET_ADDRESS, port, scriptRunner }
