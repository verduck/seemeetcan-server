**사용자 목록 조회**
---
  JSON 포맷의 사용자 정보 데이터 반환
* **URL**
  * /prfile
* **Method**
  * `POST`
* **URL Params**
  * 없음
* **Data Params**
  `{ Id: (int) }`
* **Success Response**
  * **Code:** 200<br />
  * **Content:** `{ Id: (int), StudentID: (string), Name: (string), Gender: (bool), Age: (int), Height: (int), MBTI: (string), FavoriteId ([]int) }`
* **Error Response**
  * **Code:** 400<br />
  * **Content:** `-`<br />
Or
  * **Code:** 500<br />
  * **Content:** `-`