# -*- mode: Python -*-

load('ext://restart_process', 'docker_build_with_restart')
load('ext://helm_resource', 'helm_resource', 'helm_repo')
update_settings(max_parallel_updates=10, k8s_upsert_timeout_secs=180)

# Spin up firestore
helm_resource(
  'gcp-firestore-emulator',
  'local-dependency/firestore/helm',
  port_forwards=["8081:8081"],
  labels='firestore',
)