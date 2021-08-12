package controllers

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/RianWardanaPutra/phonebook-api/httputil"
	"github.com/RianWardanaPutra/phonebook-api/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// ListContacts godoc
// @Summary Show list of contacts
// @Description get string by ID
// @Tags contacts
// @Param q query string false "Contact Name search"
// @Produce  json
// @Success 200 {array} models.Contact
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /contacts [get]
func (c *Controller) ListContacts(ctx *gin.Context) {
	q := ctx.Query("q")

	contacts, err := models.AllContacts(c.db, q)

	if err != nil {
		httputil.NewError(ctx, http.StatusNotFound, err)
		return
	}

	ctx.JSON(http.StatusOK, contacts)
}

// AddContact godoc
// @Summary Create new contact
// @Description create new contact in database, throw error if record exists
// @Tags contacts
// @Accept  json
// @Param contact body models.Contact true "Contact body"
// @Produce  json
// @Success 200 {object} models.Contact
// @Failure 400 {object} httputil.HTTPError
// @Failure 409 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /contacts [post]
func (c *Controller) AddContact(ctx *gin.Context) {
	reqBody, _ := ioutil.ReadAll(ctx.Request.Body)
	var contact models.Contact

	json.Unmarshal(reqBody, &contact)

	if name := contact.Name; name == "" {
		er := errors.New("name cannot be empty")
		httputil.NewError(ctx, 400, er)
		return
	}

	if number := contact.PhoneNumber; number == "" {
		er := errors.New("number cannot be empty")
		httputil.NewError(ctx, 400, er)
		return
	}

	c.db.Create(&contact)

	if contact.Id == 0 {
		er := errors.New("records with that name already exists")
		httputil.NewError(ctx, 409, er)
		return
	}
	ctx.JSON(http.StatusCreated, contact)
}

// FindContactById godoc
// @Summary Find a contact
// @Description Find a contact by param id
// @Tags contacts
// @Param id path int true "Contact Id"
// @Produce  json
// @Success 200 {object} models.Contact
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /contacts/{id} [get]
func (c *Controller) FindContactById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		errorMessage := errors.New("bad request")
		httputil.NewError(ctx, 400, errorMessage)
		return
	}

	contact := models.GetById(c.db, id)
	if contact.Id == 0 {
		httputil.NewError(ctx, 404, gorm.ErrRecordNotFound)
		return
	}
	ctx.JSON(http.StatusOK, contact)
}

// UpdateContactById godoc
// @Summary Update contact records
// @Description find contact in database by id and update them
// @Tags contacts
// @Param id path int true "Contact Id"
// @Param contact body models.Contact true "Contact body"
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Contact
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /contacts/{id} [put]
func (c *Controller) UpdateContactById(ctx *gin.Context) {
	reqBody, _ := ioutil.ReadAll(ctx.Request.Body)
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		errorMessage := errors.New("bad request")
		httputil.NewError(ctx, 400, errorMessage)
		return
	}

	var newContact models.Contact
	json.Unmarshal(reqBody, &newContact)

	contact := models.GetById(c.db, id)
	if contact.Id == 0 {
		httputil.NewError(ctx, 404, gorm.ErrRecordNotFound)
		return
	}

	c.db.Model(&contact).Updates(&newContact)
	ctx.JSON(http.StatusOK, contact)
}

// DeleteContactById godoc
// @Summary Delete contact records
// @Description find contact in database by id and delete them
// @Tags contacts
// @Param id path int true "Contact Id"
// @Produce  json
// @Success 200 {object} models.Contact
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /contacts/{id} [delete]
func (c *Controller) DeleteContactById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		errorMessage := errors.New("bad request")
		httputil.NewError(ctx, 400, errorMessage)
		return
	}

	contact := models.GetById(c.db, id)
	if contact.Id == 0 {
		httputil.NewError(ctx, 404, gorm.ErrRecordNotFound)
		return
	}

	c.db.Delete(&contact)
	ctx.JSON(http.StatusOK, contact)
}
