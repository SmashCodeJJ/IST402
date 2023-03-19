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

key = 7
enc = encrypt(key, message)
print("%d, %s" % (key, enc))

dec = decrypt(key, enc)
print("%d, %s" % (key, dec))