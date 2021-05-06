package example1

import (
	gc "gopkg.in/check.v1"
)

/*
Fixtures are available by using one or more of the following methods in a test suite:

func (s *SuiteType) SetUpSuite(c *C) - Run once when the suite starts running.
func (s *SuiteType) SetUpTest(c *C) - Run before each test or benchmark starts running.
func (s *SuiteType) TearDownTest(c *C) - Run after each test or benchmark runs.
func (s *SuiteType) TearDownSuite(c *C) - Run once after all tests or benchmarks have finished running.
*/

//to run only this suite we can use: go test -check.f SuiteBase

type SuiteBase struct {
	Encryptable Blob
}

// SetUpSuite this is a fixture:
func (s *SuiteBase) SetUpSuite(c *gc.C) {
	s.Encryptable.Data = []byte("this text will be encrypted")

}
func (s *SuiteBase) TestEncrypt(c *gc.C) {
	sTest := s.Encryptable.Encrypt([]byte("qwsdfrt4tghyt5fg"))
	c.Assert(sTest, gc.NotNil)
	s.Encryptable.CipherText = sTest
	message := s.Encryptable.Decryption([]byte("qwsdfrt4tghyt5fg"))
	c.Assert(message, gc.NotNil)
	c.Assert(string(message), gc.Equals, string(s.Encryptable.Data), gc.Commentf("not matched"))
}
