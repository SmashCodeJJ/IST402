alphabet = 'abcdefghijklmnopqrstuvwxyz'
ALPHABET = 'ABCDEFGHIJKLMNOPQRSTUVWXYZ'

def encrypt(n, plaintext):
    result = ''

    for l in plaintext:
        try:
            if l.isupper():
                index = ALPHABET.index(l)

                i = (index + n) % 26

                result += ALPHABET[i]
            else:
                index = alphabet.index(l)
                i = (index + n) % 26

                result += alphabet[i]
        except ValueError:
            result += l
    return result

def decrypt(n, ciphertext):
    result = ''
    for l in ciphertext:
        try:
            if l.isupper():
                index = ALPHABET.index(l)

                i = (index - n) % 26
                result += ALPHABET[i]
            else:
                index = alphabet.index(l)
                i = (index - n) % 26
                result += alphabet[i]
        except ValueError:
            result += l
    return result

message = "Encryption Is An Interesting Topic"


for i in range(26):
    enc = encrypt(i, message)
    print("%d, %s" % (i, enc))

# dec = decrypt(key, enc)
# print("%d, %s" % (key, dec))