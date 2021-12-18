**사용자 목록 조회**
---
  좋아요와 좋아요 취소
* **URL**
  * /list
* **Method**
  * `POST`
* **URL Params**
  * 없음
* **Data Params**
  `{ 'UserId': (int), 'FavoriteId': (int) }`
* **Success Response**
  * **Code:** 200<br />
  * **Content:** `{ 'result': (int) }`
* **Error Response**
  * **Code:** 400<br />
  * **Content:** `-`<br />
Or
  * **Code:** 500<br />
  * **Content:** `-`