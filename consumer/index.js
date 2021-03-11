// handleContract handles a contract
const handleContract = async (contract) => {
    let name = undefined

    try {
        name = contract.name
    } catch (_) {
        return false
    }

    if (name === undefined) return false
    return true
}

module.exports = { handleContract }
