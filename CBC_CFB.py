from Crypto.Cipher import AES
from Crypto.Util.Padding import pad, unpad
from Crypto.Random import get_random_bytes

#b'\xfd\xf0!\xa3\xf6o\x83J\xd7\xc6\xc0E)\x9b;\xca'

# Function to perform CBC encryption and decryption
def cbc_mode(key, plaintext):
    iv = get_random_bytes(AES.block_size) # generate a random initialization vector

    # creates a new AES cipher object in CBC
    # using the AES (Advanced Encryption Standard) in CBC mode
    cipher_encrypt = AES.new(key, AES.MODE_CBC, iv=iv)
    #print('ci',cipher)

    # using padding for multiple of the AES block size
    # encode plaintext into byte
    ciphertext = cipher_encrypt.encrypt(pad(plaintext.encode(), AES.block_size))
    #print(ciphertext)

    cipher_decrypt = AES.new(key, AES.MODE_CBC, iv=iv)
    decryptedtext = unpad(cipher_decrypt.decrypt(ciphertext), AES.block_size).decode()
    #print(decryptedtext)
    return iv, ciphertext.hex(), decryptedtext

# Function to perform CFB encryption and decryption
def cfb_mode(key, plaintext):
    iv = get_random_bytes(AES.block_size) # generate a random initialization vector

    # using AES in CFB mode
    cipher_encrypt = AES.new(key, AES.MODE_CFB, iv=iv, segment_size=8*AES.block_size)
    ciphertext = cipher_encrypt.encrypt(plaintext.encode())

    cipher_decrypt = AES.new(key, AES.MODE_CFB, iv=iv, segment_size=8 * AES.block_size)
    decryptedtext = cipher_decrypt.decrypt(ciphertext).decode()
    return iv, ciphertext.hex(), decryptedtext

# Accept plaintext input from user
plaintext = input("Enter plaintext: ")

# Generate a random 256-bit (32 byte)key
# in order to secure the plaintext
key = get_random_bytes(32)

# Perform CBC encryption and decryption
cbc_iv, cbc_ciphertext, cbc_decryptedtext = cbc_mode(key, plaintext)

# Perform CFB encryption and decryption
cfb_iv, cfb_ciphertext, cfb_decryptedtext = cfb_mode(key, plaintext)

# Display details about the process
print("Key:", key.hex())
print("CBC mode:")
print("  IV:", cbc_iv.hex())
print("  Ciphertext:", cbc_ciphertext)
print("  Decrypted text:", cbc_decryptedtext)
print("CFB mode:")
print("  IV:", cfb_iv.hex())
print("  Ciphertext:", cfb_ciphertext)
print("  Decrypted text:", cfb_decryptedtext)
