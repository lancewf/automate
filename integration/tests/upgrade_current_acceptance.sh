test_name="upgrade_current_acceptance"
test_upgrades=true
test_upgrade_inspec_profiles=()
test_diagnostics_filters="~skip-for-deep-upgrade"
test_upgrade_begin_manifest="current.json"

do_prepare_upgrade() {
    set_test_manifest "acceptance.json"
}
