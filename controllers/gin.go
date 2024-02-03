package controllers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"monitoriong.wiki/trackid-first-backend/database"
)

func CreateTrackid(c *gin.Context) {
	var Trackid *database.Trackid
	err := c.ShouldBind(&Trackid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	res := database.DB.Create(Trackid)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "error creating a Trackid",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Trackid": Trackid,
	})
	return
}

func ReadTrackid(c *gin.Context) {
	var Trackid database.Trackid
	id := c.Param("id")
	res := database.DB.Find(&Trackid, id)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Trackid not found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Trackid": Trackid,
	})
	return
}

func ReadTid(c *gin.Context) {
	var Trackid database.Trackid
	tidq := c.Param("tid")
	res := database.DB.Select("SELECT *")
	fmt.Println(&Trackid, tidq)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "TrackNumid not found", "tidq": tidq,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Trackid": Trackid,
	})
	return
}

func ReadTrackids(c *gin.Context) {
	var Trackids []database.Trackid
	res := database.DB.Find(&Trackids)
	if res.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": errors.New("authors not found"),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Trackids": Trackids,
	})
	return
}

func UpdateTrackid(c *gin.Context) {
	var Trackid database.Trackid
	id := c.Param("id")
	err := c.ShouldBind(&Trackid)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	var updateTrackid database.Trackid
	res := database.DB.Model(&updateTrackid).Where("id = ?", id).Updates(Trackid)

	if res.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Trackid not updated",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Trackid": Trackid,
	})
	return
}

func DeleteTrackid(c *gin.Context) {
	var Trackid database.Trackid
	id := c.Param("id")
	res := database.DB.Find(&Trackid, id)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Trackid not found",
		})
		return
	}
	database.DB.Delete(&Trackid)
	c.JSON(http.StatusOK, gin.H{
		"message": "Trackid deleted successfully",
	})
	return
}
