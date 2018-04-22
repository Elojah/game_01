# game_01

<mxfile userAgent="Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Ubuntu Chromium/65.0.3325.181 Chrome/65.0.3325.181 Safari/537.36" version="8.5.12" editor="www.draw.io" type="device"><diagram id="ce25982a-664e-2d64-380f-91dff441ab05" name="Page-1">zVpRk6I4EP41Pt4WSSDg46yzd7u1dVdWOVd3+5jBqNQisQKOer/+giRAEhxZJ4jzMqQhDXzd/X2d4ATNtsc/ONlt/mRLmk6gtzxO0PMEQuBDLP6VllNlCaOgMqx5spQXNYZF8h+VRk9a98mS5tqFBWNpkex0Y8yyjMaFZiOcs4N+2Yql+l13ZE0twyImqW39J1kWm8oaBV5j/0qT9UbdGXjyzCuJf64522fyfhOIVue/6vSWKF/y+nxDluzQMqEvEzTjjBXV0fY4o2mJrYKtmvf7hbP1c3OaFX0mwGrCG0n38tXjNCnnQpwKB59fuThal0fyeYuTwuj8lrT044nTh01S0MWOxOXZg8gKYdsU21SMQD37jfKCHi8+KqgBEIlF2ZYW/CQukROQhOxUY16ND02AQnXNphUcLG1E5sS69tzgIg4kNN0w+RZMZC9ueh2kK7CQfFel7yo5llB+CCcJjI9sYKIOXCIHuGALFwsBkeG78nCV0uNTWZriLWm2lIfPcUryPIl1WFZJms5YyvjZhaqgcp54vn/LjPsUqOEPmYDVjenSKmwDN8EkhK9p0QqsDWULuqADOmXjNCVF8qbfsQtPeYc5S8rSUhntGxkdGhHJ2Z7HVE5qF7Hhp2Yf6QgHhqPqjS1H5+jWb90r4IEV8K8vL/PFxajvOItpnl8niGXCRSUkLBPjjPESfRekAXRkat5tBRh3BBhAB8URWVjl8SlNyfPnHsShABQhozwj6aJg/Cxa92fasC/TumAUAK5TyjX6oMekaLGEGP1QZxoC8YBvUsh5MKc8EU9N+S/zSlWummC0qSZ6IKoJwI1UY/jB08GYBtitScF+0szKBr5h29d9D4pxoK0Bvp+2ArvpGFZc66qBCLXr5pNXj90UBxyzEqAhltBM4b6lYDqq+64BasGW3b+f525FV7y2M9ENR1RdYPekT/NvFlb3bMzVBHwhYe5BJqEDWW3EU5NOQRAAXZLPFquAaYtVPDXJGatEtuRWuTAW04Ap1MXSDGTv/j4EmqNgQNW1m9YRVVdNCIzu5Y4qrCjJUeEgHGpdp9DW6OOd58O0mc76TDgN36+dCykv4kBOrct25QX5Ow88xZ0P3GRH5fHWeoL2cuZp9n3UaoJjVhOy4IgZpxYeI6izH46nzh3bKhYkQ7T6QBPkQKMqnaaCm2VZ9WNtjqrS4FFIyr91380HVxy50+WOHei/nl4c7Lu5IJTIgMG3K6du7V3vzEMXjW3To76zYYQRNqQbBh+Xbm1djOxSGXfXSH22u9R49q6Ua62ww0qxO1gSF4ICx9Rc39Tc4H7igmzNzU9Z/Aiai9F4mot6bK/9QmPf6uPlmvgyOzR8E8GwRTi/CTeKgAZbE6NRtRerbPzomhgbomOtNNwxCurRnfWXmKm2uVqGHED3MR8rvqHREoXejfE1HQ0ZX3vHcHzFqKVWvX9HUzUYNdo91diLVhyOCIfdUVhYDPMhpt2NerreRKGpNyC8pDe38Qjq6EbhqNrh692o9cGlt3ZgfT+1XhAOwC1TK3ce+csNNpbGsKPMXH25EcPmd24VtM2PCdGX/wE=</diagram></mxfile>

## How to start

```
docker-compose -d
make dep
make api && bin/game_api bin/config_api.json
make client && bin/game_client bin/config_client.json
```

## TODO
- ip service (+ token association)
- ack service
- handler controller
    + resolve TODOs
- `tile38` Actor Service
- `storage/actor.go` to group use_cases
- HTTPS service for users with token creation based on PG named `auth`
- Response server to update all clients with delta compression named `sync`
- Edit `client` to make it sensitive to `sync` calls and save in a local *rocksdb ?*
