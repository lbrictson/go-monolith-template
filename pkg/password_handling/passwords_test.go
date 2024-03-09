package password_handling

import "testing"

func TestComparePassword(t *testing.T) {
	type args struct {
		hashedPassword string
		password       string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test ComparePassword",
			args: args{
				hashedPassword: HashAndSaltPassword("password"),
				password:       "password",
			},
			want: true,
		},
		{
			name: "Test ComparePassword",
			args: args{
				hashedPassword: HashAndSaltPassword("password"),
				password:       "password1",
			},
			want: false,
		},
		{
			name: "Test ComparePassword",
			args: args{
				hashedPassword: HashAndSaltPassword("Ashsh8845345!!!$#@^3468sd8gasdgSDASDH46346346FDSHFDHDFHDFH"),
				password:       "password1",
			},
			want: false,
		},
		{
			name: "Test ComparePassword",
			args: args{
				hashedPassword: HashAndSaltPassword("Ashsh8845345!!!$#@^3468sd8gasdgSDASDH46346346FDSHFDHDFHDFH"),
				password:       "Ashsh8845345!!!$#@^3468sd8gasdgSDASDH46346346FDSHFDHDFHDFH",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ComparePassword(tt.args.hashedPassword, tt.args.password); got != tt.want {
				t.Errorf("ComparePassword() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGenerateRandomPassword(t *testing.T) {
	firstRandomPassword := GenerateRandomPassword(10)
	type args struct {
		length int
	}
	tests := []struct {
		name     string
		args     args
		dontWant string
	}{
		{
			name: "Test GenerateRandomPassword",
			args: args{
				length: 10,
			},
			dontWant: firstRandomPassword,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenerateRandomPassword(tt.args.length); got == tt.dontWant {
				t.Errorf("GenerateRandomPassword() = %v, got identical result which is not expected %v", got, tt.dontWant)
			}
		})
	}
}

func TestIsPasswordValid(t *testing.T) {
	type args struct {
		password         string
		minLength        int
		complexityNeeded bool
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test IsPasswordValid: password is too short",
			args: args{
				password:         "password",
				minLength:        8,
				complexityNeeded: true,
			},
			want: false,
		},
		{
			name: "Test IsPasswordValid: password needs numbers",
			args: args{
				password:         "passwordAABB!!",
				minLength:        8,
				complexityNeeded: true,
			},
			want: false,
		},
		{
			name: "Test IsPasswordValid: password needs uppers",
			args: args{
				password:         "passwordddddd11111!",
				minLength:        8,
				complexityNeeded: true,
			},
			want: false,
		},
		{
			name: "Test IsPasswordValid: password needs lowers",
			args: args{
				password:         "PASSSSSSWOOOORDDDD121212!!",
				minLength:        8,
				complexityNeeded: true,
			},
			want: false,
		},
		{
			name: "Test IsPasswordValid: valid",
			args: args{
				password:         "passwordAABB!!23124124SADsda897shash%%%@#%gdsagsdg",
				minLength:        8,
				complexityNeeded: true,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPasswordValid(tt.args.password, tt.args.minLength, tt.args.complexityNeeded); got != tt.want {
				t.Errorf("IsPasswordValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
