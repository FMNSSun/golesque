package golesque_test

import (
	. "golesque"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"io/ioutil"
	"fmt"
)

var _ = Describe("Golesque", func() {
	Context("Foobar", func() {
		It("Can run tests", func() {
			files, _ := ioutil.ReadDir("./bintests")
			for _, f := range files {
				fmt.Printf("[%s]\n",f.Name())
            	data, err := ioutil.ReadFile("./bintests/" + f.Name())

				if err != nil {
					fmt.Println(err.Error())
				}

				Dump(data)

				context := &GLSQContext{
					Sp:    0,
					Stack: make([]*GLSQObj, 32),
				}

				err = Run(data, context)

				Expect(err).ToNot(HaveOccurred())

				fmt.Println("\n")
			}
		})

		
	})
})
