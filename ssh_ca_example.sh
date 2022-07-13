#!/usr/bin/env bash


hosts='
'

users='
'


gen_ssh_keypair() {
    local comment="${1}"
    local private_key_filename="${2}"

    if [ -z "${comment}" ]; then
        echo 'Undefined comment.'
        return 1
    fi

    if [ -z "${private_key_filename}" ]; then
        echo 'Undefined private key filename.'
        return 2
    fi

    if [ -e "${private_key_filename}" ]; then
        echo 'Private key already exists.'
        return 3
    else
        ssh-keygen                       \
            -a 2147483647                \
            -C "${comment}"              \
            -f "${private_key_filename}" \
            -N ''                        \
            -t ed25519
    fi

    # -a, rounds for key derivation function (max value 2^31-1)
    # -C, comment
    # -f, private key filename (append .pub to this for public key filename)
    # -N, symmetric passphrase
    # -t, key type (-b bitsize option ignored for this key type)
}


sign_ssh_pubkey() {
    local type="${1}"
    local identity="${2}"
    local principals="${3}"
    local key_to_sign="${4}"
    local signing_key="${5}"
    local serial_num="${6}"

    if [ -z "${type}" ]; then
        echo 'Undefined type.'
        return 1
    fi

    if [ -z "${identity}" ]; then
        echo 'Undefined identity.'
        return 2
    fi

    if [ -z "${principals}" ]; then
        echo 'Undefined principals.'
        return 3
    fi

    if [ -z "${key_to_sign}" ]; then
        echo 'Undefined key to sign.'
        return 4
    fi

    if [ -z "${signing_key}" ]; then
        echo 'Undefined signing key.'
        return 5
    fi

    if [ 'host' == "${type}" ]; then
        ssh-keygen              \
            -h                  \
            -I "${identity}"    \
            -n "${principals}"  \
            -s "${signing_key}" \
            -V '-5m:+403d'      \
            -z "${serial_num}"  \
            "${key_to_sign}"
    elif [ 'user' == "${type}" ]; then
        ssh-keygen              \
            -I "${identity}"    \
            -n "${principals}"  \
            -s "${signing_key}" \
            -V '-5m:+403d'      \
            -z "${serial_num}"  \
            "${key_to_sign}"
    else
        echo 'Undefined type.'
    fi

    # -h, create a host certificate instead of a user certificate
    # -I, identity
    # -n, principals
    # -s, signing private key
    # -V, validity interval (5 minutes ago to 1*Y+M+W+D from now)
    # -z, serial number (default 0), use '+' prefix to autoincrement
    # public key file to sign
}


main() {
    gen_ssh_keypair 'hostca' 'hostca'
    gen_ssh_keypair 'userca' 'userca'

    for host in ${hosts}; do
        gen_ssh_keypair "root@${host}" "ssh_host_ed25519_key_${host}"
        sign_ssh_pubkey 'host' "${host}" "${host}" "ssh_host_ed25519_key_${host}" 'hostca' 0
    done

    for user in ${users}; do
        gen_ssh_keypair "${user}" "id_ed25519_${user}"
        sign_ssh_pubkey 'user' "${user}" "${user}" "id_ed25519_${user}" 'userca' 0
    done
}


main


# For hosts, add to /etc/ssh/sshd_config on each host:
#   HostCertificate /etc/ssh/new_host_key-cert.pub
#   (uncomment HostKey /etc/ssh/new_host_key)
# Add to ~/.ssh/known_hosts:
#   @cert-authority * $(cat hostsca.pub)

# For users,...
