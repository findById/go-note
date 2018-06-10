---------------------
| len 4 |    data   |
---------------------
|0|0|0|0|    ...    |
---------------------
```java
public void test() throws IOException {
    byte[] bytes = encode("a".getBytes());
    byte[] result = decode(bytes);
    System.out.println(new String(result));
}

private byte[] encode(byte[] data) throws IOException {
    ByteArrayOutputStream os = new ByteArrayOutputStream();
    int len = data.length;
    os.write((len >> 24) & 0xFF);
    os.write((len >> 16) & 0xFF);
    os.write((len >> 8) & 0xFF);
    os.write((len) & 0xFF);
    os.write(data);
    return os.toByteArray();
}

private byte[] decode(byte[] data) throws IOException {
    ByteArrayInputStream is = new ByteArrayInputStream(data);
    byte[] buf = new byte[4];
    is.read(buf);
    int len = ((buf[0] & 0xFF) << 24) | ((buf[1] & 0xFF) << 16) | ((buf[2] & 0xFF) << 8) | (buf[3] & 0xFF);

    byte[] bytes = new byte[len];
    is.read(bytes);
    return bytes;
}
```
