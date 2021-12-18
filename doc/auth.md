**사용자 조회**
---
  학번으로 JSON 포맷의 사용자 데이터 반환
* **URL**
  * /
* **Method**
  * `POST`
* **URL Params**
  * 없음
* **Data Params**
  `{ StudentID: (string) }`
* **Success Response**
  * **Code:** 200<br />
  * **Content:** `{ Id: (int), StudentID: (string), Name: (string), Gender: (bool), Age: (int), Height: (int), MBTI: (string), FavoriteId ([]int) }`
* **Error Response**
  * **Code:** 400<br />
  * **Content:** `-`
  Or
  * **Code:** 500<br />
  * **Content:** `-`