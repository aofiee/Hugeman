# Hugeman

## 1. pre-commit install
```bash
pre-commit install --install-hooks
```

## 2. dependency check
```bash
brew install dependency-check
```

## 3. config dependency check in cmd/
```bash
echo '<?xml version="1.0" encoding="UTF-8"?>\n<suppressions xmlns="https://jeremylong.github.io/DependencyCheck/dependency-suppression.1.3.xsd">\n</suppressions>' > dependency-suppressions.xml
```

## 4. 1st scan
```bash
dependency-check --out cmd/ --scan cmd/ --project $(basename $PWD)/cmd --enableExperimental --failOnCVSS 0 --suppression cmd/dependency-suppressions.xml
```