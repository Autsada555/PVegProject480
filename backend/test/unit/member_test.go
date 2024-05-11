package unit

import (
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
	"github.com/Autsada55/P5VegProject480/entity"
)

func TestMemberfail(t *testing.T) {
	g := NewGomegaWithT(t)
	t.Run(`ต้องมากกว่า 7 แต่น้อยกว่า 100 ตัว`, func(t *testing.T) {
		member := entity.Customer{
			Firstname:      "John",
			Lastname:       "John",
			Email:          "john@example.com",
			Password:       "46", //password wrong
		}

		ok, err := govalidator.ValidateStruct(member)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).To(ContainSubstring("Password ต้องมากกว่า 7 แต่น้อยกว่า 100 ตัว"))
	})

	t.Run(`Email is invalid`, func(t *testing.T) {
		member := entity.Customer{
			Firstname: "John",
			Lastname:  "John",
			Email:     "john", //email wrong
			Password:  "12345678",
		}

		ok, err := govalidator.ValidateStruct(member)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).To(ContainSubstring("Email is invalid"))
	})

}

func TestAllMember(t *testing.T) {
	g := NewGomegaWithT(t)
	t.Run(`All correct`, func(t *testing.T) {
		member := entity.Customer{
			Firstname:      "John",
			Lastname:       "John",
			Email:          "john@example.com",
			Password:       "46213456",
		}
		ok, err := govalidator.ValidateStruct(member)

		g.Expect(ok).To(BeTrue())
		g.Expect(err).To(BeNil())
	})

}
