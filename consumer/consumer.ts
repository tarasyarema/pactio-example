// Contract is the consumer defined contract that it expects
type Contract = {
    name: string;
    email: string;
}

// handleContract handles a contract
const handleContract = async (contract: Contract): Promise<boolean> => {
    try {
        const { name, email } = contract

        if (name === undefined ||
            email === undefined) {
            throw "Bad contract"
        }
    } catch (_) {
        return false
    }

    return true
}

export { Contract, handleContract }
