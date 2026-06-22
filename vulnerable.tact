contract Vulnerable {
    owner: Address;

    receive(msg: Message) {
        // TACT-001: Missing seqno check
        // TACT-002: Missing sender check
        selfdestruct();
        var x: Int;
        x = x + 1; // TACT-004
        while(true) { // TACT-006
            // infinite loop
        }
    }
}
