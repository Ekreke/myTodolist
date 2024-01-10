package service

type UserGetProjectsService struct {
}

// func UserGetProjects(token string) serializer.Response {
// 	db := conf.DB
// 	code := e.SUCCESS
// 	var username string

// 	// can't get token
// 	if token == "" {
// 		code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
// 		return serializer.Response{
// 			Status: code,
// 			Msg:    e.GetMsg(code),
// 		}
// 	} else {
// 		// get claims
// 		claims, err := utils.ParseUserToken(token)
// 		if err != nil {
// 			logging.Info(err)
// 			code = e.ERROR_AUTH_TOKEN
// 			return serializer.Response{
// 				Status: code,
// 				Msg:    e.GetMsg(code),
// 			}
// 		} else if time.Now().Unix() > claims.ExpiresAt {
// 			code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
// 			return serializer.Response{
// 				Status: code,
// 				Msg:    e.GetMsg(code),
// 			}
// 		} else {
// 			username = claims.Username
// 		}
// 	}
// 	var projectsIds []model.Project
// 	db.Debug().Where("username = ?", username).Select("projects_id").Find(&projects)
// 	db.Debug().Where("projects_id =?", projectsIds)
// }
