# an array with 4 rows and 2 columns
codebook = [[0b00, 0b01], [0b01, 0b10], [0b10, 0b11], [0b11, 0b00]]
message = [0b00, 0b01, 0b10, 0b11]
iv = 0b10

def codebookLookup(xor):
    lookupValue = 0
    for i in range(4):
        if codebook[i][0] == xor:
            lookupValue = codebook[i][1]
            break
    return lookupValue

def codebookReverseLookup(lookupValue):
    reverseValue = 0
    for i in range(4):
        if codebook[i][1] == lookupValue:
            reverseValue = codebook[i][0]
            break
    return reverseValue

x = 0
feedback = iv
for i in range(4):
    xor = message[x] ^ feedback
    lookupValue = codebookLookup(xor)
    feedback = codebookReverseLookup(lookupValue)
    x += 1
    print(f"The ciphered value of {message[i]:b} is {lookupValue:b}")

x = 0
feedback = iv
for i in range(4):
    xor = lookupValue ^ feedback
    reverseValue = codebookReverseLookup(xor)
    feedback = message[x]
    x += 1
    print(f"The decrypted value of {message[i]:b} is {reverseValue:b}")
