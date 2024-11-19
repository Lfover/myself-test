def parse_response(res, file):
print("--------------------------- response ---------------------------")
# print(f"response raw bytes: {res}")
protocol_version = res[0] >> 4
header_size = res[0] & 0x0f
message_type = res[1] >> 4
message_type_specific_flags = res[1] & 0x0f
serialization_method = res[2] >> 4
message_compression = res[2] & 0x0f
reserved = res[3]
header_extensions = res[4:header_size*4]
payload = res[header_size*4:]
print(f"            Protocol version: {protocol_version:#x} - version {protocol_version}")
print(f"                 Header size: {header_size:#x} - {header_size * 4} bytes ")
print(f"                Message type: {message_type:#x} - {MESSAGE_TYPES[message_type]}")
print(f" Message type specific flags: {message_type_specific_flags:#x} - {MESSAGE_TYPE_SPECIFIC_FLAGS[message_type_specific_flags]}")
print(f"Message serialization method: {serialization_method:#x} - {MESSAGE_SERIALIZATION_METHODS[serialization_method]}")
print(f"         Message compression: {message_compression:#x} - {MESSAGE_COMPRESSIONS[message_compression]}")
print(f"                    Reserved: {reserved:#04x}")
if header_size != 1:
print(f"           Header extensions: {header_extensions}")
if message_type == 0xb:  # audio-only server response
if message_type_specific_flags == 0:  # no sequence number as ACK
print("                Payload size: 0")
return False
else:
sequence_number = int.from_bytes(payload[:4], "big", signed=True)
payload_size = int.from_bytes(payload[4:8], "big", signed=False)
payload = payload[8:]
print(f"             Sequence number: {sequence_number}")
print(f"                Payload size: {payload_size} bytes")
file.write(payload)
if sequence_number < 0:
return True
else:
return False

elif message_type == 0xf:
code = int.from_bytes(payload[:4], "big", signed=False)
msg_size = int.from_bytes(payload[4:8], "big", signed=False)
error_msg = payload[8:]
if message_compression == 1:
error_msg = gzip.decompress(error_msg)
error_msg = str(error_msg, "utf-8")
print(f"          Error message code: {code}")
print(f"          Error message size: {msg_size} bytes")
print(f"               Error message: {error_msg}")
return True
elif message_type == 0xc:
msg_size = int.from_bytes(payload[:4], "big", signed=False)
payload = payload[4:]
if message_compression == 1:
payload = gzip.decompress(payload)
print(f"            Frontend message: {payload}")
else:
print("undefined message type!")
return True