# Simple gitcloak Initialization and Encrypting Single File

In this example, a single file `demo.txt` is encrypted using `gitcloak`.

* testing with demo.txt

```bash
echo "Hello World" > demo.txt
echo ""           >> demo.txt
echo "Bye World"  >> demo.txt
echo ""           >> demo.txt
gitcloak init -e aes -k passwordpassword -s 123456 -p demo.txt
```

* The above command will create file `.gitcloak/config.yaml` with following content:
```yaml
rules:
    - name: gcinit
      encryption:
        algorithm: aes
        key: passwordpassword
        seed: 123456
      path: demo.txt
```
* More rules can be added in the above configuration file.
* Let us encrypt the content of `demo.txt` file
```bash
gitcloak encrypt
```

* Now contents of `demo.txt` is encrypted
```bash
$ cat demo.txt 
6pc9s09/zHHm4NFGlozJ8K0tMCwzI8USqXJ/
kjvgI/519kpGVKvZO3IHDg==
YUEpdAdvRWZXmVd345+bq+JKb6d0g/oe7w==
kjvgI/519kpGVKvZO3IHDg==
```

* Once, you are ensured that the file is encrypted, it can be committed to git.
```bash
git add demo.txt
git commit -am "gitcloak demo commit"
```

* Run following command for encrypting, if somefile is left decrypted, and rollback feature to be implemented.
```bash
gitcloak commit -m "Added demo.txt in gitcloak"
```

* Decrypting `demo.txt` for reading/updation
```bash
gitcloak decrypt
```

* And it reveals the contents of the file
```bash
$ cat demo.txt 
Hello World

Bye World

```

### Updating the encrypted file
* Updating the file `demo.txt`

* First decrypt the file
```bash
gitcloak decrypt
```

* Append *Hello World* in `demo.txt` file
```bash
vi demo.txt 
```

* Again Encrypt the files
```bash
gitcloak encrypt
```

* Now commit the encrypted updated `demo.txt` file
```bash
git commit -am "Updated: appended line Hello World in demo.txt"
```
* Now `demo.txt` is encrypted, and please notice lines #1, and #5
```bash
$ cat demo.txt 
6pc9s09/zHHm4NFGlozJ8K0tMCwzI8USqXJ/
kjvgI/519kpGVKvZO3IHDg==
YUEpdAdvRWZXmVd345+bq+JKb6d0g/oe7w==
kjvgI/519kpGVKvZO3IHDg==
6pc9s09/zHHm4NFGlozJ8K0tMCwzI8USqXJ/
```