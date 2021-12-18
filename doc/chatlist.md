**사용자 목록 조회**
---
  JSON 포맷의 채팅방 목록 데이터 반환
* **URL**
  * /chatlist
* **Method**
  * `POST`
* **URL Params**
  * 없음
* **Data Params**
  `{ 'UserId': (int) }`
* **Success Response**
  * **Code:** 200<br />
  * **Content:** `{ 'List': ([]object) { 'Id': (int), 'UUID': ([]byte]), 'Name': (string), 'LastMessage': (string), 'Time': (timestamp) } }`
* **Error Response**
  * **Code:** 400<br />
  * **Content:** `-`<br />
Or
  * **Code:** 500<br />
  * **Content:** `-`