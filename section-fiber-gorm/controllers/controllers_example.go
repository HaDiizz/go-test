package controllers

import (
	"errors"
	"regexp"
	"strconv"
	"strings"

	"github.com/HaDiizz/database"
	m "github.com/HaDiizz/models"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func HelloTest(c *fiber.Ctx) error {
	return c.SendString("Hello World")
}

func Factorial(c *fiber.Ctx) error {
	numParam := c.Params("num")
	num, err := strconv.Atoi(numParam)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	result, err := calculateFactorial(num)

	if err != nil {
		// return c.SendStatus(fiber.StatusBadRequest)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"result": result})
}

func calculateFactorial(num int) (int, error) {
	if num < 0 {
		return 0, errors.New("non-negative num")
	}
	if num == 0 || num == 1 {
		return 1, nil
	}

	result, err := calculateFactorial(num - 1)
	if err != nil {
		return 0, err
	}

	return num * result, nil
}

func AsciiGenerator(c *fiber.Ctx) error {
	text := c.Query("tax_id")

	store := []rune{}
	result := ""
	for _, v := range text {

		store = append(store, v)
	}

	for i, v := range store {

		result += strconv.Itoa(int(v))

		if i < len(store)-1 {
			result += " "
		}
	}

	return c.JSON(result)

}

func Register(c *fiber.Ctx) error {

	user := new(m.Registration)
	if err := c.BodyParser((user)); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})

	}

	validate := validator.New()
	validate.RegisterValidation("username_validate", func(fl validator.FieldLevel) bool {
		return regexp.MustCompile(`^[a-zA-Z0-9_-]+$`).MatchString(fl.Field().String())
	})
	validate.RegisterValidation("web_validate", func(fl validator.FieldLevel) bool {
		return regexp.MustCompile(`^[a-z0-9\.]+$`).MatchString(fl.Field().String())
	})
	err := validate.Struct(user)
	if err != nil {
		fieldErrors := make(map[string]string)

		for _, e := range err.(validator.ValidationErrors) {
			if e.Field() == "Email" && e.Tag() == "email" {
				fieldErrors[strings.ToLower(e.Field())] = "Invalid email"
			} else if e.Field() == "Username" && e.Tag() == "username_validate" {
				fieldErrors[strings.ToLower(e.Field())] = "ใช้อักษรภาษาอังกฤษ (a-z), (A-Z), ตัวเลข (0-9) และเครื่องหมาย (_), (-) เท่านั้น เช่น Example_01"
			} else if e.Field() == "Password" && (e.Tag() == "min" || e.Tag() == "max") {
				fieldErrors[strings.ToLower(e.Field())] = "ความยาว 6-20 อักษร"
			} else if e.Field() == "WebName" && (e.Tag() == "min" || e.Tag() == "max") {
				fieldErrors[strings.ToLower(e.Field())] = "ความยาว 2-30 อักษร"
			} else if e.Field() == "WebName" && (e.Tag() == "web_validate") {
				fieldErrors[strings.ToLower(e.Field())] = "ใช้อักษรภาษาอังกฤษตัวเล็ก (a-z), ตัวเลข (0-9) ห้ามใช้เครื่องหมายอักขระพิเศษยกเว้นขีด (-) ห้ามเว้นวรรค และห้ามใช้ภาษาไทย"
			} else {
				fieldErrors[strings.ToLower(e.Field())] = e.Field() + " is required"
			}
		}

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Validation errors occurred",
			"errors":  fieldErrors,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"result": user})
}

// CRUD dogs

func GetDeletedDogs(c *fiber.Ctx) error {
	db := database.DBConn
	var dogs []m.Dogs

	db.Unscoped().Where("deleted_at is NOT NULL").Find(&dogs)
	return c.Status(200).JSON(dogs)
}

func GetDogsRangeCountByDogId(c *fiber.Ctx) error {
	db := database.DBConn
	var dogs []m.Dogs

	db.Unscoped().Where("dog_id > ? && dog_id < ?", 50, 100).Find(&dogs)
	return c.Status(200).JSON(dogs)
}

func GetDogs(c *fiber.Ctx) error {
	db := database.DBConn
	var dogs []m.Dogs

	db.Find(&dogs) //delete = null
	return c.Status(200).JSON(dogs)
}

func GetDog(c *fiber.Ctx) error {
	db := database.DBConn
	search := strings.TrimSpace(c.Query("search"))
	var dog []m.Dogs

	result := db.Find(&dog, "dog_id = ?", search)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}
	return c.Status(200).JSON(&dog)
}

func AddDog(c *fiber.Ctx) error {
	db := database.DBConn
	var dog m.Dogs

	if err := c.BodyParser(&dog); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	db.Create(&dog)
	return c.Status(201).JSON(dog)
}

func UpdateDog(c *fiber.Ctx) error {
	db := database.DBConn
	var dog m.Dogs
	id := c.Params("id")

	if err := c.BodyParser(&dog); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	db.Where("id = ?", id).Updates(&dog)
	return c.Status(200).JSON(dog)
}

func RemoveDog(c *fiber.Ctx) error {
	db := database.DBConn
	id := c.Params("id")
	var dog m.Dogs
	result := db.Delete(&dog, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.SendStatus(200)
}

func GetDogsJsonSummary(c *fiber.Ctx) error {
	db := database.DBConn
	var dogs []m.Dogs

	db.Find(&dogs)

	sumRed := 0
	sumGreen := 0
	sumPink := 0
	sumNone := 0

	var dataResults []m.DogsRes
	for _, v := range dogs {
		typeStr := ""
		if v.DogID >= 10 && v.DogID <= 50 {
			typeStr = "red"
			sumRed++
		} else if v.DogID >= 100 && v.DogID <= 150 {
			typeStr = "green"
			sumGreen++
		} else if v.DogID >= 200 && v.DogID <= 250 {
			typeStr = "pink"
			sumPink++
		} else {
			typeStr = "no color"
			sumNone++
		}

		d := m.DogsRes{
			Name:  v.Name,
			DogID: v.DogID,
			Type:  typeStr,
		}
		dataResults = append(dataResults, d)
	}

	r := m.ResultData{
		Data:        dataResults,
		Name:        "golang-test",
		Count:       len(dogs),
		Sum_red:     sumRed,
		Sum_green:   sumGreen,
		Sum_pink:    sumPink,
		Sum_noColor: sumNone,
	}
	return c.Status(200).JSON(r)
}

func GetDogsJson(c *fiber.Ctx) error {
	db := database.DBConn
	var dogs []m.Dogs

	db.Find(&dogs)

	var dataResults []m.DogsRes
	for _, v := range dogs {
		typeStr := ""
		if v.DogID == 111 {
			typeStr = "red"
		} else if v.DogID == 113 {
			typeStr = "green"
		} else if v.DogID == 999 {
			typeStr = "pink"
		} else {
			typeStr = "no color"
		}

		d := m.DogsRes{
			Name:  v.Name,
			DogID: v.DogID,
			Type:  typeStr,
		}
		dataResults = append(dataResults, d)
	}

	type ResultData struct {
		Data  []m.DogsRes `json:"data"`
		Name  string      `json:"name"`
		Count int         `json:"count"`
	}
	r := ResultData{
		Data:  dataResults,
		Name:  "golang-test",
		Count: len(dogs), //หาผลรวม,
	}
	return c.Status(200).JSON(r)
}

// CRUD companies

func GetCompanies(c *fiber.Ctx) error {
	db := database.DBConn
	var companies []m.Companies

	db.Find(&companies) //delelete = null
	return c.Status(200).JSON(companies)
}

func GetCompany(c *fiber.Ctx) error {
	db := database.DBConn
	search := strings.TrimSpace(c.Query("search"))
	var company []m.Companies

	result := db.Find(&company, "id = ?", search)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}
	return c.Status(200).JSON(&company)
}

func AddCompany(c *fiber.Ctx) error {
	db := database.DBConn
	var company m.Companies

	if err := c.BodyParser(&company); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	db.Create(&company)
	return c.Status(201).JSON(company)
}

func UpdateCompany(c *fiber.Ctx) error {
	db := database.DBConn
	var company m.Companies
	id := c.Params("id")

	if err := c.BodyParser(&company); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	db.Where("id = ?", id).Updates(&company)
	return c.Status(200).JSON(company)
}

func RemoveCompany(c *fiber.Ctx) error {
	db := database.DBConn
	id := c.Params("id")
	var company m.Companies
	result := db.Delete(&company, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.SendStatus(200)
}

// CRUD profiles

func GetProfiles(c *fiber.Ctx) error {
	db := database.DBConn
	var profiles []m.UserProfiles

	db.Find(&profiles) //delete = null
	return c.Status(200).JSON(profiles)
}

func GetProfile(c *fiber.Ctx) error {
	db := database.DBConn
	search := strings.TrimSpace(c.Query("search"))
	var profile []m.UserProfiles

	result := db.Find(&profile, "employee_id = ?", search)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}
	return c.Status(200).JSON(&profile)
}

func AddProfile(c *fiber.Ctx) error {
	db := database.DBConn
	var profile m.UserProfiles

	if err := c.BodyParser(&profile); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	var existingEmpId m.UserProfiles

	if err := db.Where("employee_id = ?", profile.EmployeeID).First(&existingEmpId).Error; err == nil {
		if existingEmpId.EmployeeID == profile.EmployeeID {
			return c.Status(fiber.StatusConflict).JSON(fiber.Map{"message": "Employee ID already exists."})
		}
	}

	validate := validator.New()

	err := validate.Struct(profile)
	if err != nil {
		fieldErrors := make(map[string]string)

		for _, e := range err.(validator.ValidationErrors) {
			if e.Field() == "Email" && e.Tag() == "email" {
				fieldErrors[strings.ToLower(e.Field())] = "Invalid email"
			} else if e.Field() == "Age" && e.Tag() == "min" {
				fieldErrors[strings.ToLower(e.Field())] = "Age must be greater than or equal 18"
			} else {
				fieldErrors[strings.ToLower(e.Field())] = e.Field() + " is required"
			}
		}

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Validation errors occurred",
			"errors":  fieldErrors,
		})
	}

	db.Create(&profile)
	return c.Status(201).JSON(profile)
}

func UpdateProfile(c *fiber.Ctx) error {
	db := database.DBConn
	var profile m.UserProfiles
	id := c.Params("id")

	if err := c.BodyParser(&profile); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	if profile.EmployeeID != "" {
		var existingEmpId m.UserProfiles
		if err := db.Where("employee_id = ?", profile.EmployeeID).First(&existingEmpId).Error; err == nil {
			if existingEmpId.EmployeeID == profile.EmployeeID {
				return c.Status(fiber.StatusConflict).JSON(fiber.Map{"message": "Employee ID already exists."})
			}
		}
	}

	db.Where("id = ?", id).Updates(&profile)
	return c.Status(200).JSON(profile)
}

func RemoveProfile(c *fiber.Ctx) error {
	db := database.DBConn
	id := c.Params("id")
	var profile m.UserProfiles
	result := db.Delete(&profile, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.SendStatus(200)
}

func GetProfileAnyAges(c *fiber.Ctx) error {
	var profiles []m.UserProfiles
	db := database.DBConn

	sumGenZ := 0
	sumGenY := 0
	sumGenX := 0
	sumBabyBoomer := 0
	sumGI := 0

	db.Find(&profiles)
	var dataResult []m.UserProfileResult
	_ = dataResult
	for _, value := range profiles {

		genStr := ""

		if value.Age > 75 {
			genStr = "G.I Generation"
			sumGI++
		} else if value.Age >= 24 && value.Age <= 41 {
			genStr = "Gen Y"
			sumGenY++

		} else if value.Age >= 42 && value.Age <= 56 {
			genStr = "Gen X"
			sumGenX++
		} else if value.Age >= 57 && value.Age <= 75 {
			genStr = "Baby Boomer"
			sumBabyBoomer++
		} else {
			genStr = "Gen Z"
			sumGenZ++
		}

		p := m.UserProfileResult{
			EmployeeID: value.EmployeeID,
			Name:       value.Name,
			LastName:   value.LastName,
			Birthday:   value.Birthday,
			Age:        value.Age,
			Email:      value.Email,
			Tel:        value.Tel,
			Gen:        genStr,
		}

		dataResult = append(dataResult, p)

	}

	r := m.UserProfileAgesResult{
		Data:          dataResult,
		Name:          "profile-ages",
		Count:         len(profiles),
		SumGenZ:       sumGenZ,
		SumGenY:       sumGenY,
		SumGenX:       sumGenX,
		SumBabyBoomer: sumBabyBoomer,
		SumGI:         sumGI,
	}

	return c.Status(200).JSON(r)

}

func SearchData(c *fiber.Ctx) error {

	db := database.DBConn
	var profiles []m.UserProfiles

	search := c.Query("search")

	db.Where("(employee_id = ? || name = ? || last_name = ? ) && deleted_at is NULL", search, search, search).Find(&profiles)

	return c.JSON(profiles)

}
