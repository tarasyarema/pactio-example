// Contract is the consumer defined contract that it expects
type Contract = {
    name: string;
    email: string;
    address: Address;
}

// Address is a new field added in v2
type Address = {
    street: string;
    number: number;
}

// handleContract handles a contract
const handleContract = async (contract: Contract): Promise<boolean> => {
    try {
        const { name, email, address: { street, number } } = contract
        console.log(name, email, street, number)
    } catch (_) {
        return false
    }

    return true
}

export { Contract, Address, handleContract }
