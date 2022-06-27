# Order Application
Birbirleri ile haberleşen minimum iki microservice'in olduğu bir yapıdır. Müşteri (customer) ve sipariş (order) oluşturulması hedeflenmiştir. 

## Çalıştırma (Running) ve Kullanımı (How to use)
`docker compose up` : Bu komutu order-application içerisinde çalıştırarak projeyi ayağa kaldırabilirsiniz.

`docker compose down` : container'ları silmek için kullanabilirsiniz. Image'lar silinmez.

API Base Endpoint (BASE_ENPOINT): `http://localhost:5000/`

### Customer Endpoints

`[HttpGet] [BASE_ENDPOINT]/customer`

`[HttpGet] [BASE_ENDPOINT]/customer/{id}`

`[HttpGet] [BASE_ENDPOINT]/customer/validate/{id}`

`[HttpPost] [BASE_ENDPOINT]/customer`

`[HttpPut] [BASE_ENDPOINT]/customer`

`[HttpDelete] [BASE_ENDPOINT]/customer/{id}`


- **Örnekler curl ile (Examples with curl)**

Müşteri listeleme (Customer List): `curl -X 'GET' \ '[BASE_ENDPOINT]/customer?page=1&limit=10' \ -H 'accept: text/plain'`

Müşteri oluşturma (Create Customer): `curl -X 'POST' \
  '[BASE_ENDPOINT]/customer' \
  -H 'accept: text/plain' \
  -H 'Content-Type: application/json' \
  -d '{
    "name":"NameTest",
    "email": "name@test.com",
    "address": {
        "addressLine":"AdresTest",
        "country": "CountryTest",
        "city":"CityTest",
        "cityCode": 34
    }
}'`

Müşteri detayları (Customer Details): `curl -X 'GET' \
  '[BASE_ENDPOINT]/customer/{id}' \
  -H 'accept: text/plain'
`

Müşteri sil (Delete Customer): `curl -X 'DELETE' \
  '[BASE_ENDPOINT]/customer/{id}' \
  -H 'accept: */*'
`

### Order Endpoints

`[HttpGET] [BASE_ENDPOINT]/order`

`[HttpGET] [BASE_ENDPOINT]/order/{id}`

`[HttpGET] [BASE_ENDPOINT]/order/GetByCustomerId/{id}`

`[HttpPost] [BASE_ENDPOINT]/order`

`[HttpPut] [BASE_ENDPOINT]/order`

`[HttpPut] [BASE_ENDPOINT]/order/{id}`

`[HttpDelete] [BASE_ENDPOINT]/order/{id}`


- **Örnekler curl ile (Examples with curl)**

Order listeleme (Order List): `curl -X 'GET' \ '[BASE_ENDPOINT]/order?page=1&limit=10' \ -H 'accept: text/plain'`

Order Oluşturma (Create Order): `curl -X 'POST' \ '[BASE_ENDPOINT]/order' \ -H 'accept: text/plain' \ -d ''`

Order listeleme (Order List): `curl -X 'GET' \ '[BASE_ENDPOINT]/order/{id}' \ -H 'accept: text/plain'`

## Teknik Tasarım
#### Müşteriler
Sistemde teorik anlamda sınırsız sayıda müşteri kaydı yapılabilmektedir.

Veri yapısındaki alanlar aşağıdaki gibidir:
- UUID
- Ad (Name)
- Mail (Email)
- Adres Bilgisi (Address)
  - Adres (AddressLine)
  - Şehir (City)
  - Ülke (Country)
  - Şehir Kodu (CityCode)

#### Siparişler
Sistemde teorik anlamda sınırsız sayıda sipariş kaydı yapılabilmektedir.

Veri yapısındaki alanlar aşağıdaki gibidir:
- UUID
- Müşteri UUID
- Miktar (Quantity)
- Fiyat (Price)
- Statü (Status)
- Adres Bilgisi (Address)
  - Adres (AddressLine)
  - Şehir (City)
  - Ülke (Country)
  - Şehir Kodu (CityCode)
- Ürün (Product)
  - UUID
  - Resim Url (ImageUrl)
  - İsim (Name)


## Teknik Detaylar
- Kullanılan teknolojiler:
   - Golang
   - .NET Core
   - MongoDB
   - Ocelot
   - Docker compose

## Mimari Yapı
![orderApp](https://user-images.githubusercontent.com/79265067/175916391-69353e6d-eefc-4b6e-b198-fb039be8faf8.png)
