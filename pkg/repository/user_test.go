package repository

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-playground/assert/v2"
	"gokul.go/pkg/utils/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Test_User_UserSignUp(t *testing.T) {

	type args struct {
		input    models.UserDetails
		referral string
	}
	tests := []struct {
		name    string
		args    args
		stub    func(sqlmock.Sqlmock)
		want    models.UserDeatilsResponse
		wantErr error
	}{
		{
			name: "success signup user",
			args: args{input: models.UserDetails{
				Name:            "sidhu",
				Email:           "sidhu@gmail.com",
				Phone:           "6282288363",
				Password:        "1234",
				ConfirmPassWord: "1234",
			}, referral: "7BDH"},

			stub: func(mockSQL sqlmock.Sqlmock) {
				expectedQuery := `INSERT INTO users(.+)$`
				mockSQL.ExpectQuery(expectedQuery).WithArgs("sidhu", "sidhu@gmail.com", "1234", "6282288363", "7BDH").
					WillReturnRows(sqlmock.NewRows([]string{"id", "name", "email", "phone"}).AddRow(1, "sidhu", "sidhu@gmail.com", "6282288363"))

			},
			want: models.UserDeatilsResponse{
				Id:    1,
				Name:  "sidhu",
				Email: "sidhu@gmail.com",
				Phone: "6282288363"},
			wantErr: nil,
		},
		{
			name: "error signup user",
			args: args{
				input: models.UserDetails{
					Name:            "",
					Email:           "",
					Phone:           "",
					Password:        "",
					ConfirmPassWord: "",
				}, referral: "",
			},
			stub: func(mockSQL sqlmock.Sqlmock) {
				expectedQuery := `INSERT INTO users(.+)$`
				mockSQL.ExpectQuery(expectedQuery).WithArgs("", "", "", "", "").
					WillReturnError(errors.New("text string"))

			},
			want:    models.UserDeatilsResponse{},
			wantErr: errors.New("text string"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockDB, mockSQL, _ := sqlmock.New()
			defer mockDB.Close()

			gormDB, _ := gorm.Open(postgres.New(postgres.Config{
				Conn: mockDB,
			}), &gorm.Config{})
			tt.stub(mockSQL)
			u := NewUserRepository(gormDB)
			got, err := u.UserSignUp(tt.args.input, tt.args.referral)
			assert.Equal(t, tt.wantErr, err)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userRepo.UserSIgnUp()= %v, want %v ", got, tt.want)
			}
		})
	}

}

func Test_User_CheckUserAvailability(t *testing.T) {
	tests := []struct {
		name string
		args string
		stub func(sqlmock.Sqlmock)
		want bool
	}{
		{
			name: "user available",
			args: "gokul@gmail.com",
			stub: func(mockSQL sqlmock.Sqlmock) {
				expectedQuery := "SELECT COUNT(*) FROM users WHERE email=$1"
				mockSQL.ExpectQuery(regexp.QuoteMeta(expectedQuery)).
					WithArgs("gokul@gmail.com").
					WillReturnRows(sqlmock.NewRows([]string{"count(*)"}).AddRow(1))
			},
			want: true,
		},
		{
			name: "user not available",
			args: "gokul@gmail.com",
			stub: func(mockSQL sqlmock.Sqlmock) {
				expectedQuery := "SELECT COUNT(*) FROM users WHERE email=$1"
				mockSQL.ExpectQuery(regexp.QuoteMeta(expectedQuery)).
					WithArgs("gokul@gmail.com").
					WillReturnRows(sqlmock.NewRows([]string{"count(*)"}).AddRow(0))
			},
			want: false,
		},
		{
			name: "error from database",
			args: "gokul@gmail.com",
			stub: func(mockSQL sqlmock.Sqlmock) {
				expectedQuery := "SELECT COUNT(*) FROM users WHERE email=$1"
				mockSQL.ExpectQuery(regexp.QuoteMeta(expectedQuery)).
					WithArgs("gokul@gmail.com").
					WillReturnError(errors.New("database error"))
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockDB, mockSQL, _ := sqlmock.New()
			defer mockDB.Close()

			gormDB, _ := gorm.Open(postgres.New(postgres.Config{
				Conn: mockDB,
			}), &gorm.Config{})

			tt.stub(mockSQL)
			u := NewUserRepository(gormDB)

			result := u.CheckUserAvailability(tt.args)
			assert.Equal(t, tt.want, result)
		})
	}
}

func Test_User_FindUserByEmail(t *testing.T) {

	tests := []struct {
		name    string
		args    models.UserLoign
		stub    func(sqlmock.Sqlmock)
		want    models.UserSignInResponse
		wantErr error
	}{
		{
			name: "finding user by email",
			args: models.UserLoign{
				Email:    "gokul@gmail.com",
				PassWord: "1234",
			},
			stub: func(mockSQL sqlmock.Sqlmock) {
				expectedQuery := "SELECT * FROM users WHERE email =$1 and blocked = false"
				mockSQL.ExpectQuery(regexp.QuoteMeta(expectedQuery)).WithArgs("gokul@gmail.com").
					WillReturnRows(sqlmock.NewRows([]string{"id", "name", "email", "phone", "password"}).AddRow(1, "gokul", "gokul@gmail.com", "6282288363", "1234"))
				fmt.Println("SQL Query:", expectedQuery)
				fmt.Println("Values:", "gokul@gmail.com")

			},
			want: models.UserSignInResponse{
				Id:       1,
				Name:     "gokul",
				Email:    "gokul@gmail.com",
				Phone:    "6282288363",
				Password: "1234",
			},
			wantErr: nil,
		},
		{
			name: "error finding user by email",
			args: models.UserLoign{
				Email:    "gokul@gmail.com",
				PassWord: "1234",
			},
			stub: func(mockSQL sqlmock.Sqlmock) {
				expectedQuery := "SELECT * FROM users WHERE email =$1 and blocked = false"
				mockSQL.ExpectQuery(regexp.QuoteMeta(expectedQuery)).WithArgs("").
					WillReturnError(errors.New("error checking user details"))
			},
			want:    models.UserSignInResponse{},
			wantErr: errors.New("error checking user details"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockDB, mockSQL, _ := sqlmock.New()
			defer mockDB.Close()

			gormDB, _ := gorm.Open(postgres.New(postgres.Config{
				Conn: mockDB,
			}), &gorm.Config{})

			tt.stub(mockSQL)

			u := NewUserRepository(gormDB)

			result, err := u.FindUserByEmail(tt.args)

			assert.Equal(t, tt.wantErr, err)
			assert.Equal(t, tt.want, result)
		})
	}
}

func Test_UserBlockStatu(t *testing.T) {
	tests := []struct {
		name    string
		args    string
		stub    func(sqlmock.Sqlmock)
		want    bool
		wantErr error
	}{
		{
			name: "cheking useres existance ",
			args: "gokul@gmail.com",
			stub: func(mockSQL sqlmock.Sqlmock) {
				expextedQuery := "SELECT blocked FROM users WHERE email = $1"
				mockSQL.ExpectQuery(regexp.QuoteMeta(expextedQuery)).WithArgs("gokul@gmail.com").
					WillReturnRows(sqlmock.NewRows([]string{"status"}).AddRow(1))
			},
			want:    true,
			wantErr: nil,
		}, {
			name: "user dont exist ",
			args: "gokul@gmail.com",
			stub: func(mockSQL sqlmock.Sqlmock) {
				expextedQuery := "SELECT blocked FROM users WHERE email = $1"
				mockSQL.ExpectQuery(regexp.QuoteMeta(expextedQuery)).WithArgs("gokul@gmail.com").
					WillReturnRows(sqlmock.NewRows([]string{"status"}))
			},
			want:    false,
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockDB, mockSQL, _ := sqlmock.New()
			defer mockDB.Close()

			gormDB, _ := gorm.Open(postgres.New(postgres.Config{
				Conn: mockDB,
			}), &gorm.Config{})

			tt.stub(mockSQL)

			u := NewUserRepository(gormDB)

			result, err := u.UserBlockStatus(tt.args)

			assert.Equal(t, tt.wantErr, err)
			assert.Equal(t, tt.want, result)
		})
	}
}

func Test_AddAddress(t *testing.T) {

	type args struct {
		id    int
		input models.AddAddress
	}

	tests := []struct {
		name    string
		args    args
		stub    func(sqlmock.Sqlmock, args)
		wantErr error
	}{
		{
			name: "adding the address",
			args: args{input: models.AddAddress{
				Name:      "gokul",
				HouseName: "vettiyankal",
				Street:    "west",
				City:      "newyork",
				State:     "kerala",
				Pin:       "685565",
			}, id: 1},

			stub: func(sqlMock sqlmock.Sqlmock, arg args) {
				expectedQuery := `(?i)INSERT INTO addresses\(users_id ,name ,house_name,street,city,state,pin\)\s*VALUES\(\$1, \$2, \$3, \$4 ,\$5, \$6, \$7\)`
				sqlMock.ExpectExec(expectedQuery).
					WithArgs(1, "gokul", "vettiyankal", "west", "newyork", "kerala", "685565").
					WillReturnResult(sqlmock.NewResult(1, 1))

			},
			wantErr: nil,
		},
		{
			name: "cant add the addresses",
			args: args{input: models.AddAddress{
				Name:      " ",
				HouseName: " ",
				Street:    " ",
				City:      " ",
				State:     " ",
				Pin:       " ",
			}, id: 1},
			stub: func(sqlmock sqlmock.Sqlmock, arg args) {
				expectedQuery := `(?i)INSERT INTO addresses\(users_id ,name ,house_name,street,city,state,pin\)\s*VALUES\(\$1, \$2, \$3, \$4 ,\$5, \$6, \$7\)`
				sqlmock.ExpectExec(expectedQuery).
					WithArgs(1, " ", " ", " ", " ", " ", " ").
					WillReturnError(errors.New("error adding address"))

			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			fmt.Println("name", tt.name)
			mockDB, mockSQL, _ := sqlmock.New()
			defer mockDB.Close()

			gormDB, _ := gorm.Open(postgres.New(postgres.Config{
				Conn: mockDB,
			}), &gorm.Config{})

			tt.stub(mockSQL, tt.args)
			u := NewUserRepository(gormDB)

			err := u.AddAddress(tt.args.id, tt.args.input)
			fmt.Println("actual error is ", err)

			assert.Equal(t, tt.wantErr, err)
		})
	}

}
