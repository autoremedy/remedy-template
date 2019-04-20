# Auto Remedy Template

Pull down the template
```
$ faas template pull https://github.com/autoremedy/remedy-template
$ faas new --list
Languages available as templates:
- go-remedy
```

Creating a new remedy
```
$ faas new --lang go-remedy <remedy-name>
```

```
$ faas build -f <remedy-name>.yml
```

```
faas deploy -f <remedy-name>.yml --env=combine_output=false
```

Note that if you are developing on minikube, you need to change

```
kubectl edit -n openfaas-fn deployment.apps/<remedy-name>
```
change `imagePullPolicy` from `Always` to `Never`

```
echo '{"status":"ok"}' | faas invoke <remedy-name>
```
