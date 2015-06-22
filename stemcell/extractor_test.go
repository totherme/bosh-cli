package stemcell_test

import (
	"errors"

	. "github.com/cloudfoundry/bosh-init/stemcell"
	. "github.com/cloudfoundry/bosh-init/internal/github.com/onsi/ginkgo"
	. "github.com/cloudfoundry/bosh-init/internal/github.com/onsi/gomega"

	biproperty "github.com/cloudfoundry/bosh-init/internal/github.com/cloudfoundry/bosh-utils/property"
	fakesys "github.com/cloudfoundry/bosh-init/internal/github.com/cloudfoundry/bosh-utils/system/fakes"

	fakebistemcell "github.com/cloudfoundry/bosh-init/stemcell/fakes"
)

var _ = Describe("Manager", func() {
	var (
		extractor             Extractor
		fs                    *fakesys.FakeFileSystem
		reader                *fakebistemcell.FakeStemcellReader
		stemcellTarballPath   string
		stemcellExtractionDir string

		expectedExtractedStemcell ExtractedStemcell
	)

	BeforeEach(func() {
		fs = fakesys.NewFakeFileSystem()
		reader = fakebistemcell.NewFakeReader()
		stemcellTarballPath = "/stemcell/tarball/path"
		stemcellExtractionDir = "/path/to/dest"
		fs.TempDirDir = stemcellExtractionDir

		extractor = NewExtractor(reader, fs)

		expectedExtractedStemcell = NewExtractedStemcell(
			Manifest{
				Name:      "fake-stemcell-name",
				ImagePath: "fake-image-path",
				CloudProperties: biproperty.Map{
					"fake-prop-key": "fake-prop-value",
				},
			},
			stemcellExtractionDir,
			fs,
		)
		reader.SetReadBehavior(stemcellTarballPath, stemcellExtractionDir, expectedExtractedStemcell, nil)
	})

	Describe("Extract", func() {
		Context("when the extraction succeeeds", func() {
			It("extracts and parses the stemcell manifest", func() {
				stemcell, err := extractor.Extract(stemcellTarballPath)
				Expect(err).ToNot(HaveOccurred())
				Expect(stemcell).To(Equal(expectedExtractedStemcell))

				Expect(reader.ReadInputs).To(Equal([]fakebistemcell.ReadInput{
					{
						StemcellTarballPath: stemcellTarballPath,
						DestPath:            stemcellExtractionDir,
					},
				}))
			})

			It("does not delete the destination file path", func() {
				extractor.Extract(stemcellTarballPath)
				Expect(fs.FileExists(stemcellExtractionDir)).To(BeTrue())
			})
		})

		Context("when the read fails", func() {
			BeforeEach(func() {
				reader.SetReadBehavior(stemcellTarballPath, stemcellExtractionDir, expectedExtractedStemcell, errors.New("fake-read-error"))
			})

			It("returns an error", func() {
				_, err := extractor.Extract(stemcellTarballPath)
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("fake-read-error"))
			})

			It("deletes the destination file path", func() {
				extractor.Extract(stemcellTarballPath)
				Expect(fs.FileExists(stemcellExtractionDir)).To(BeFalse())
			})
		})
	})
})
