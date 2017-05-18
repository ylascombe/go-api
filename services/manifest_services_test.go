package services

import (
	"testing"
	"github.com/stretchr/testify/assert"
)


func TestBuildCommand(t *testing.T) {
	commands := BuildCommands("LOCAL.ini", "", "fake_path")

	assert.Equal(t, 10, len(commands))
	assert.Equal(t, "ansible-playbook -i inventories/LOCAL.ini plateforme_reactive.yml", commands[3])

	assert.Equal(t, "ansible-playbook -i inventories/LOCAL.ini exploit_download_nexus_artifact.yml -e artifact_group=fr.laposte.colis.colis360 -e artifact=spark-colis-360 -e repository=releases -e version=1.4.2", commands[4])
	assert.Equal(t, "ansible-playbook -i inventories/LOCAL.ini deploy_spark_app.yml -e spark_app_name=spark-colis-360 -e spark_app_version=1.4.2 -e spark_app_filename=spark-colis-360-1.4.2 -assembly.jar -e force_deploy=true --vault-password-file fake_path", commands[5])
	assert.Equal(t, "ansible-playbook -i inventories/LOCAL.ini deploy_apiserver.yml -e apiserver_instance=api-colis-360 -e apiserver_version=1.4.3 --vault-password-file fake_path", commands[6])

	assert.Equal(t, "ansible-playbook -i inventories/LOCAL.ini exploit_download_nexus_artifact.yml -e artifact_group=fr.laposte.colis.datasafe -e artifact=spark-datasafe -e repository=releases -e version=0.0.7", commands[7])
	assert.Equal(t, "ansible-playbook -i inventories/LOCAL.ini deploy_spark_app.yml -e spark_app_name=spark-datasafe -e spark_app_version=0.0.7 -e spark_app_filename=spark-datasafe-0.0.7 -assembly.jar -e force_deploy=true --vault-password-file fake_path", commands[8])
	assert.Equal(t, "ansible-playbook -i inventories/LOCAL.ini deploy_apiserver.yml -e apiserver_instance=api-datasafe -e apiserver_version=0.0.8 --vault-password-file fake_path", commands[9])

}
