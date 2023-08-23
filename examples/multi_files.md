# gitcloak Initialization and Encrypting Multiple Files

In this example, multiple files (wildcard `demo*.txt`) matching regex `demo.*.txt$` in are encrypted using `gitcloak`.

## More Example Patterns

| Files Matching Description | Example Wildcard | regex  `-r` option | Remarks |
| -------------------------- | ---------------- | ------------------ | ------- |
| Single matching file in multiple directories | `.env` | `.env` | `-p` / `--path` can be used |
| Single matching file in single directory | `/prod/.env` | `/prod/.env` | `-p` / `--path` can be used |
| Multiple matching files in single directory | `/prod/.env*` | `/prod/.env.*$` | |
| Multiple matching files in multiple matching directories | `/*apps*/.env*` | `/.*apps.*/.env.*$` | |


----

* testing with demo.txt

```bash
echo "Hello World" > demo.txt
echo ""           >> demo.txt
echo "Bye World"  >> demo.txt
echo ""           >> demo.txt
echo "Hello World" > demo2.txt
echo ""           >> demo2.txt
echo "Bye World"  >> demo2.txt
echo ""           >> demo2.txt
gitcloak init -e aes -k passwordpassword -s 234567 -r 'demo.*.txt$'
```

* Output
```
2023/08/23 13:01:37 ...t/gitcloak/cmd/init.go:34 ...k/cmd.glob..func7 - [INFO]  gitcloak init started
2023/08/23 13:01:37 ...t/gitcloak/cmd/init.go:261 ...eateGitCommitHook - [INFO]  Create git hook /tmp/test2/.git/hooks/pre-commit
2023/08/23 13:01:37 ...t/gitcloak/cmd/init.go:83 ...k/cmd.glob..func7 - [INFO]  gitcloak Commit Hash = a09ce1c5addbde2400d2740d4b6ef44c7c0367ac
2023/08/23 13:01:37 ...t/gitcloak/cmd/init.go:88 ...k/cmd.glob..func7 - [INFO]  git Commit Hash = 69a41eeaad0959e857dea4b75214e8641e4b31d0
2023/08/23 13:01:37 ...t/gitcloak/cmd/init.go:101 ...k/cmd.glob..func7 - [INFO]  gitcloak init completed with aes algorithm
```


* The above command will create file `.gitcloak/config.yaml` with following content:
```yaml
rules:
    - name: gcinit
      encryption:
        algorithm: aes
        key: passwordpassword
        seed: 234567
      path_regex: demo.*.txt$
```
* More rules can be added in the above configuration file.

* Let us encrypt the content of `demo.txt` file
```bash
gitcloak encrypt
```
* Output
```
2023/08/23 13:01:54 ...itcloak/cmd/encrypt.go:21 ...k/cmd.glob..func6 - [INFO]  gitcloak encrypt started
2023/08/23 13:01:54 ...itcloak/cmd/encrypt.go:31 ...k/cmd.glob..func6 - [INFO]  Processing Rule : 0
2023/08/23 13:01:54 .../pkg/encrypt/common.go:137 ...rypt.EncryptFiles - [INFO]  Encrypting File: /tmp/test2/demo.txt
2023/08/23 13:01:54 .../pkg/encrypt/common.go:137 ...rypt.EncryptFiles - [INFO]  Encrypting File: /tmp/test2/demo2.txt
2023/08/23 13:01:54 .../pkg/encrypt/common.go:263 ...RuleForEncryption - [INFO]  No error
```

* Now contents of `demo*.txt` is encrypted
```bash
$ head demo*.txt
==> demo2.txt <==
r5aR3Ea29nNvtm/J5vUO+mE1TZFAhGP8y+g3
/h/f3fCMDq302wpl9KiHaw==
qWB2TJLGYp3vT9azbmPs7lyTunitlfQTMQ==
/h/f3fCMDq302wpl9KiHaw==

==> demo.txt <==
0P+tsct2/sbkFvC1hUsAZQHbwM5kRfXR/qhu
nLwHLfED/tYYv/rbcU8pIA==
jAio7ky6p28SowDub3ywU1KcqAH2qRLpww==
nLwHLfED/tYYv/rbcU8pIA==
```

* Once, you are ensured that the file is encrypted, it can be committed to git.
```bash
git add demo*.txt
git commit -am "gitcloak demo commit"
```
* Output
```
gitcloak git pre-commit hook
2023/08/23 13:21:15 ...gitcloak/cmd/commit.go:25 ...k/cmd.glob..func2 - [INFO]  gitcloak commit called
2023/08/23 13:21:15 ...gitcloak/cmd/commit.go:40 ...k/cmd.glob..func2 - [INFO]  Processing Rule : 0
2023/08/23 13:21:15 .../pkg/encrypt/common.go:134 ...rypt.EncryptFiles - [INFO]  Encrypted already : /tmp/test2/demo.txt
2023/08/23 13:21:15 .../pkg/encrypt/common.go:134 ...rypt.EncryptFiles - [INFO]  Encrypted already : /tmp/test2/demo2.txt
2023/08/23 13:21:15 .../pkg/encrypt/common.go:263 ...RuleForEncryption - [INFO]  No error
2023/08/23 13:21:15 ...gitcloak/cmd/commit.go:90 ...md.gitcloakCommit - [INFO]  gitcloak Commit Hash = 4c6e57837bf22ce12fabc0fce0af5a1a60ca2ee6
2023/08/23 13:21:15 ...gitcloak/cmd/commit.go:95 ...md.gitcloakCommit - [INFO]  git Commit Hash = 69a41eeaad0959e857dea4b75214e8641e4b31d0
2023/08/23 13:21:15 ...gitcloak/cmd/commit.go:102 ...md.gitcloakCommit - [INFO]  gitcloak commit completed with git pre-commit hook called message
[master ce6de3b] Working example
 3 files changed, 9 insertions(+)
 create mode 100644 .gitignore
 create mode 100644 demo.txt
 create mode 100644 demo2.txt
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

* Difference, which shows that 1 line is appended
```bash
$ git diff
diff --git a/demo.txt b/demo.txt
index fefa65d..234f7e9 100644
--- a/demo.txt
+++ b/demo.txt
@@ -2,3 +2,4 @@
 nLwHLfED/tYYv/rbcU8pIA==
 jAio7ky6p28SowDub3ywU1KcqAH2qRLpww==
 nLwHLfED/tYYv/rbcU8pIA==
+0P+tsct2/sbkFvC1hUsAZQHbwM5kRfXR/qhu
```

* Now commit the encrypted updated `demo.txt` file
```bash
git commit -am "Updated: appended line Hello World in demo.txt"
```
* Now `demo.txt` is encrypted, and please notice lines #1, and #5
```bash
$ cat demo.txt 
0P+tsct2/sbkFvC1hUsAZQHbwM5kRfXR/qhu
nLwHLfED/tYYv/rbcU8pIA==
jAio7ky6p28SowDub3ywU1KcqAH2qRLpww==
nLwHLfED/tYYv/rbcU8pIA==
0P+tsct2/sbkFvC1hUsAZQHbwM5kRfXR/qhu
```
