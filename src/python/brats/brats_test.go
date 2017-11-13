package brats_test

import (
	"github.com/cloudfoundry/libbuildpack/bratshelper"
	"github.com/cloudfoundry/libbuildpack/cutlass"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"golang.org/x/crypto/bcrypt"
)

var _ = Describe("Python buildpack", func() {
	bratshelper.UnbuiltBuildpack("python", CopyBrats)
	bratshelper.DeployingAnAppWithAnUpdatedVersionOfTheSameBuildpack(CopyBrats)
	bratshelper.StagingWithBuildpackThatSetsEOL("python", func(_ string) *cutlass.App {
		return CopyBrats("2.2.x")
	})
	bratshelper.StagingWithADepThatIsNotTheLatest("python", CopyBrats)
	bratshelper.StagingWithCustomBuildpackWithCredentialsInDependencies(`python\-[\d\.]+\-linux\-x64\-[\da-f]+\.tgz`, CopyBrats)
	bratshelper.DeployAppWithExecutableProfileScript("python", CopyBrats)
	bratshelper.DeployAnAppWithSensitiveEnvironmentVariables(CopyBrats)
	bratshelper.ForAllSupportedVersions("python", CopyBrats, func(pythonVersion string, app *cutlass.App) {
		PushApp(app)

		By("installs the correct version of Python", func() {
			Expect(app.Stdout.String()).To(ContainSubstring("Installing python " + pythonVersion))
			Expect(app.GetBody("/version")).To(ContainSubstring(pythonVersion))
		})
		By("runs a simple webserver", func() {
			Expect(app.GetBody("/")).To(ContainSubstring("Hello World!"))
		})
		By("parses XML with nokogiri", func() {
			Expect(app.GetBody("/nokogiri")).To(ContainSubstring("Hello, World"))
		})
		By("supports EventMachine", func() {
			Expect(app.GetBody("/em")).To(ContainSubstring("Hello, EventMachine"))
		})
		By("encrypts with bcrypt", func() {
			hashedPassword, err := app.GetBody("/bcrypt")
			Expect(err).ToNot(HaveOccurred())
			Expect(bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte("Hello, bcrypt"))).ToNot(HaveOccurred())
		})
		By("supports bson", func() {
			Expect(app.GetBody("/bson")).To(ContainSubstring("00040000"))
		})
		By("supports postgres", func() {
			Expect(app.GetBody("/pg")).To(ContainSubstring("could not connect to server: No such file or directory"))
		})
		By("supports mysql2", func() {
			Expect(app.GetBody("/mysql2")).To(ContainSubstring("Unknown MySQL server host 'testing'"))
		})
	})
})
