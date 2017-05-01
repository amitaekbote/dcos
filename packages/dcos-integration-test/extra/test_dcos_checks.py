import uuid


def test_components(dcos_api_session):
    # target contains a dcos-checks subcommand as a key and a list of tulip of optional arguments.
    target = {
        "components": ()
    }

    cmd_tpl = "/opt/mesosphere/bin/dcos-checks --role agent {} {}"
    cmds = [cmd_tpl.format(subcommand, " ".join(arg for arg in args)) for subcommand, args in target.items()]
    test_uuid = uuid.uuid4().hex
    check_job = {
        'id': 'test-dcos-checks-' + test_uuid,
        'run': {
            'cpus': .1,
            'mem': 128,
            'disk': 0,
            'cmd': " && ".join(cmds)}}

    dcos_api_session.metronome_one_off(check_job)
