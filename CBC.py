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

x = 0
lookupValue = 0
for i in range(4):
    if x == 0:
        xor = message[x] ^ iv

    else:
        xor = message[x] ^ lookupValue
    lookupValue = codebookLookup(xor)
    x += 1
    print(f"The ciphered value of a is {lookupValue:b}")
