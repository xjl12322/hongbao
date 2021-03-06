package utils

//// JWT 签名结构
//type JWT struct {
//	SigningKey []byte
//}
//
//// 一些常量
//var (
//	TokenExpired     error  = errors.New("Token is expired")
//	TokenNotValidYet error  = errors.New("Token not active yet")
//	TokenMalformed   error  = errors.New("That's not even a token")
//	TokenInvalid     error  = errors.New("Couldn't handle this token:")
//	SignKey          string = "newtrekWang"
//)
//
//
//// 载荷，可以加一些自己需要的信息
//type CustomClaims struct {
//	ID    string `json:"userId"`
//	Name  string `json:"name"`
//	Email  string `json:"email"`
//	Password string `json:"password"`
//	jwt.StandardClaims
//}
//
////新建一个jwt实例
//var jwtSecret = []byte("SignKey")
////// 新建一个jwt实例
////func NewJWT() *JWT {
////	return &JWT{
////		[]byte("SignKey"),
////	}
////}
//
////CreateToken 生成一个token
//func GenerateToken(id,name,email, password string) (string, error) {
//	nowTime := time.Now()
//	expireTime := nowTime.Add(300 * time.Second)
//	claims := CustomClaims{
//		id,
//		name,
//		email,
//		password,
//		jwt.StandardClaims{
//			//NotBefore: int64(time.Now().Unix() - 1000), // 签名生效时间
//			//ExpiresAt: int64(time.Now().Unix() + 3600), // 过期时间 一小时
//			ExpiresAt: expireTime.Unix(),
//			Issuer:    "gin-blog",
//		},
//	}
//	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
//	token, err := tokenClaims.SignedString(jwtSecret)
//	return token, err
//}
////// CreateToken 生成一个token
////func (j *JWT) CreateToken(claims CustomClaims) (string, error) {
////	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
////	return token.SignedString(j.SigningKey)
////}
//
//// 解析Tokne
//func (j *JWT) ParseToken(tokenString string) (*CustomClaims, error) {
//	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
//		return j.SigningKey, nil
//	})
//	if err != nil {
//		if ve, ok := err.(*jwt.ValidationError); ok {
//			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
//				return nil, TokenMalformed
//			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
//				// Token is expired
//				return nil, TokenExpired
//			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
//				return nil, TokenNotValidYet
//			} else {
//				return nil, TokenInvalid
//			}
//		}
//	}
//	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
//		return claims, nil
//	}
//	return nil, TokenInvalid
//}
//
//func ParseToken(token string) (*CustomClaims, error) {
//	tokenClaims, err := jwt.ParseWithClaims(token, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
//		return jwtSecret, nil
//	})
//	if tokenClaims != nil {
//		if claims, ok := tokenClaims.Claims.(*CustomClaims); ok && tokenClaims.Valid {
//			return claims, nil
//		}
//	}
//	return nil, err
//}
//
//////更新token
//func RefreshToken(token string) (string, error) {
//	tokenClaims, err := jwt.ParseWithClaims(token, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
//		return jwtSecret, nil
//	})
//	if err != nil {
//		return "", err
//	}
//	if tokenClaims != nil {
//		if claims, ok := tokenClaims.Claims.(*CustomClaims); ok && tokenClaims.Valid {
//					jwt.TimeFunc = time.Now
//					claims.StandardClaims.ExpiresAt = time.Now().Add(300 * time.Second).Unix()
//					return GenerateToken(claims.ID,claims.Name,claims.Email,claims.Password)
//		}
//	}
//	return "", TokenInvalid
//
//}




////更新token
//func (j *JWT) RefreshToken(tokenString string) (string, error) {
//	jwt.TimeFunc = func() time.Time {
//		return time.Unix(0, 0)
//	}
//	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
//		return j.SigningKey, nil
//	})
//	if err != nil {
//		return "", err
//	}
//	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
//		jwt.TimeFunc = time.Now
//		claims.StandardClaims.ExpiresAt = time.Now().Add(1 * time.Hour).Unix()
//		return j.CreateToken(*claims)
//	}
//	return "", TokenInvalid
//}







