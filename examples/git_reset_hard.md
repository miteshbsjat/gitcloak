# Git Time Travel to Older File Versions with Different Encryption Key/Seed

This document will help in Encryption Key/Seed Rotation, as well as getting the contents of files which were encrypted with older key/seed.

**ToDO** : The steps given here will be automated in next version of `gitcloak`.

## Prerequisite for this Usecase

Please read and get acquainted with `gitcloak` by following the given documents :
* [Initialization and Single File Encryption](examples/single_file.md)
* [Initialization and Multiple Files (Regex) Encryption](examples/multi_files.md)

## Changing the Encryption Key/Seed

* Please _decrypt_ the files before changing the encryption key/seed. This will reveal the original contents of all the encrypted files.

```bash
gitcloak decrypt
```

* Now, you can change the `key` / `seed` in `.gitcloak/config.yaml` file. e.g. `key` is changed to `passwor1passwor1` 
```yaml
rules:
    - name: gcinit
      encryption:
        algorithm: aes
        key: passwor1passwor1
        seed: 123456
      path: demo.txt
```

* Editing `demo.txt` for this example
```bash
echo "Hello World" >> demo.txt
```

* Encrypt all the files before committing
```bash
gitcloak encrypt
```

* Committing the file(s) changed due to new Encryption key
```bash
git commit -am "Second aes encryption with passwor1passwor1 key rotation"
```


* Verifying the new encryption key is working
```bash
gitcloak decrypt
cat demo.txt
```
  * output
  ```bash
  Hello World

  Bye World

  Hello World
  ```

----

## Time Travel (Rolling Back to Previous Encryption Key)

In this example, we are going from newer encryption key `passwor1passwor1` to older encryption key `passwordpassword` 

* Demo git log
```bash
commit 87d1a19c4bd1e6425f03b89080ff7d79c3a2d770 (HEAD -> master)
Author: Mitesh Singh Jat <mitesh.singh.jat@gmail.com>
Date:   Mon Sep 4 16:50:32 2023 +0530

    Second aes encryption with passwor1passwor1 key rotation

commit d79bf3abdc87f74a23513d0e9f44369f5ac34573
Author: Mitesh Singh Jat <mitesh.singh.jat@gmail.com>
Date:   Mon Sep 4 16:48:30 2023 +0530

    First aes encryption with passwordpassword

commit 92870958d876fab25f082c0c929fdf8933ebc6d9
Author: Mitesh Singh Jat <mitesh.singh.jat@gmail.com>
Date:   Mon Sep 4 16:46:21 2023 +0530

    demo
```

* Let us time travel i.e. Rollback to commit "First aes encryption with passwordpassword" which corresponds to older key
```bash
git reset --hard d79bf3abdc87f74a23513d0e9f44369f5ac34573
HEAD is now at d79bf3a First aes encryption with passwordpassword
```

* Please **note** the **previous commit id** of `First aes encryption with passwordpassword` i.e. `92870958d876fab25f082c0c929fdf8933ebc6d9`

* Go to `.gitcloak/` directory and get the corresponding gitcloak commit id
```bash
cat .gitcloak/ggcmap.txt | grep 92870958d876fab25f082c0c929fdf8933ebc6d9 | cut -d= -f2-
5b31f96d2c343d66cc71b070feaf7b18d777cfd6
```
* Or you can use [textfilekv-cli tool]()
```bash
textfilekv-cli -f .gitcloak/ggcmap.txt get -k 92870958d876fab25f082c0c929fdf8933ebc6d9
5b31f96d2c343d66cc71b070feaf7b18d777cfd6
```

* After finding gitcloak commit id, git reset to `5b31f96d2c343d66cc71b070feaf7b18d777cfd6` to get corresponding version of file `.gitcloak/config.yaml`
```bash
cd .gitcloak/
git reset --hard 5b31f96d2c343d66cc71b070feaf7b18d777cfd6
```

* Verifying the config is really rolled back to the older encryption key `passwordpassword`
```yaml
cat config.yaml 
rules:
    - name: gcinit
      encryption:
        algorithm: aes
        key: passwordpassword
        seed: 123456
      path: demo.txt
```

* Let us go back to parent dir, decrypt the rolled back contents
```
cd ..
gitcloak decrypt
```

* Verifying that contents are available as those were available in rolled back version
```bash
cat demo.txt 
Hello World

Bye World

```
