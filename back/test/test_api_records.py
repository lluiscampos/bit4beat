import subprocess
import time
import os
import pathlib
import shutil

import requests
import pytest

binary_name = "bit4beat.back"
binary_path = os.path.join(pathlib.Path(__file__).parent.parent.absolute(), binary_name)
storage_path = "/tmp/dummy-dir"
root_url = "http://localhost:8080"


@pytest.fixture(scope="session")
def backend_clean(request):
    os.makedirs(storage_path, exist_ok=True)
    shutil.rmtree(storage_path)
    os.makedirs(storage_path)
    proc = subprocess.Popen([binary_path])
    time.sleep(0.1)

    yield

    proc.kill()
    proc.communicate()

    print(proc.stdout)
    print(proc.stderr)
    shutil.rmtree(storage_path)


def test_records_post_invalid(backend_clean):
    url = root_url + "/record"
    data = {"kind": "nonexistent"}
    resp = requests.post(url, json=data)
    assert resp.status_code == 400

    json = resp.json()
    assert "Error:Field validation for 'Kind' failed" in json["error"]

    assert len(os.listdir(storage_path)) == 0


def test_records_post_valid(backend_clean):
    url = root_url + "/record"

    data = {
        "id": "stuff-id",
        "kind": "ski",
        "date": "sometime",
        "place": "somewhere",
    }
    resp = requests.post(url, json=data)
    assert resp.status_code == 200, resp
    json = resp.json()
    assert json["id"] == "stuff-id"
    assert len(os.listdir(storage_path)) == 1

    data = {
        "id": "stuff-id-other",
        "kind": "ski",
        "date": "sometime",
        "place": "somewhere",
        "distance": "1234.5 km",
        "participants": ["my", "good", "friends"],
        "reference": "nowhere",
    }
    resp = requests.post(url, json=data)
    assert resp.status_code == 200, resp.text
    json = resp.json()
    assert json["id"] == "stuff-id-other"
    assert len(os.listdir(storage_path)) == 2
