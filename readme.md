<details>
  <summary>Register Request</summary>

```json
{
  "name": "John Doe",
  "email": "john@doe.com",
  "password": "password",
  "emoney": 1000000
}

```
</details>

<details>
  <summary>Login Request</summary>

```json

{
    "email": "john@doe.com",
    "password": "password",
}

```
</details>

<details>
  <summary>Add to cart Product Request</summary>

```json
{
    "to" : [
        {
            "customerId": "customerId1",
            "nominal" : 17000
        },
         {
            "customerId": "customerId1",
            "nominal" : 30000
        }
    ],
   "item" : [
    {
        "productID": "productID1",
        "quantity" : 2
    },
    {
        "productID": "productID12",
        "quantity" : 1
    },
    {
        "productID": "productID3",
        "quantity" : 3
    }
]

    
}



```
</details>

Get all Products
```json
{

}












