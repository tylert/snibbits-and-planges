#!/usr/bin/env bash

# One keyring to rule them all and, in the darkness, bind them.

# Generate a new GPG key using as many current best practices as we can.
# Try to minimize the amount of data that needs to be backed up...
# - use ED-209 keys since they are tiny and fast
# - don't just use the GPG defaults (sign,cert key + encr key and no auth key)
# - use as few uids as possible to limit how many self-signatures they need
# - combine sign and auth functions under a single key (slight compromise, for size)
# - keep the cert key separate in case we need to rekey subkeys (optionally, offline it)
# - we can derive an SSH keypair from the GPG auth key

# Tools required:  bash, awk, gpg (2.4.x or greater), git, go

# https://en.wikipedia.org/wiki/ED-209
# https://en.linuxportal.info/encyclopedia/g/gpg-key
# https://riseup.net/en/security/message-security/openpgp/gpg-best-practices
# https://goral.net.pl/post/use-gpg-for-ssh-keys
# https://raymii.org/s/articles/GPG_noninteractive_batch_sign_trust_and_send_gnupg_keys.html

if [ -z "${UIDS}" ]; then
    UIDS=('flink <2039487520934875023948572039587239587235982735928752975@yggmail>')
fi

expires='2y'

# Generate a new primary cert key with userid
gpg --batch --pinentry-mode loopback --passphrase '' --quick-gen-key "${UIDS[0]}" ed25519 cert "${expires}"

# XXX FIXME TODO  GPG is dumb because it doesn't return the new key it just made
# XXX FIXME TODO  This script is equally dumb and assumes that the first key in the list is the one we want
# XXX FIXME TODO  We should use a fresh, temporary keyboxd location to store new keys being generated instead
fprs=($(gpg --list-keys --with-keygrip --with-subkey-fingerprint --with-colons | awk -F: '/fpr:/ {print $10}'))
# grps=($(gpg --list-keys --with-keygrip --with-subkey-fingerprint --with-colons | awk -F: '/fpr:/ {print $10}'))
# Filenames ~/.gnupg/private-keys-v1.d/*.key come from the keygrip values
# ${fprs[0]}, ${fprs[1]}, ... up to ${#fprs[@]} at index n-1

# Add any additional, desired userids
for uid in "${UIDS[@]}"; do
    gpg --quick-add-uid "${fprs[0]}" "${uid}"
done
gpg --check-trustdb

# Add new subkeys
gpg --batch --pinentry-mode loopback --passphrase '' --quick-add-key "${fprs[0]}" ed25519 sign,auth "${expires}"
gpg --batch --pinentry-mode loopback --passphrase '' --quick-add-key "${fprs[0]}" cv25519 encr "${expires}"

# Set a symmetric passphrase
# gpg --command-fd 0 --pinentry-mode loopback --change-passphrase $key << END
# ... list NEW\nNEW\n or OLD\nNEW\n passphrases here...
# END

# Extract a new SSH keypair from the GPG auth key
# git clone https://github.com/pinpox/pgp2ssh ; cd pgp2ssh ; go build . ; mv pgp2ssh .. ; cd ..
# ./pgp2ssh 2>&1 | (umask 0077 && tee tmp-ssh-sec.txt)
# ... manually strip extra gunk out of the resulting file that isn't an SSH private key...
# gpg --export-ssh-key ${key}  # verify we get the same result below
# umask 0022 ; ssh-keygen -f tmp-ssh-sec.txt -y > tmp-ssh-pub.txt
# ssh-keygen -f tmp-ssh-sec.txt -p  # set a symmetric passphrase on the private key

# Paperkey only backs up private keys which doesn't work so great if you aren't using a public key server

# Back up the actual key parts of the keys (private AND public)
# umask 0077 ; gpg --armor --export-secret-keys "${fprs[0]}" > tmp-gpg-sec.txt
# umask 0077 ; gpg --armor --export-secret-subkeys "${fprs[0]}" > tmp-gpg-ssb.txt
# gpg --armor --export-secret-keys "${fprs[0]}" | gpgsplit --prefix 'tmp-gpg-sec-'
# gpg --armor --export-secret-subkeys "${fprs[0]}" | gpgsplit --prefix 'tmp-gpg-ssb-'

# Install https://github.com/kazu-yamamoto/pgpdump
# for pkt in tmp-gpg-{sec,ssb}-000*; do pgpdump ${pkt} ; done

# XXX FIXME TODO  Replace this entire script with a Go tool using github.com/ProtonMail/gopenpgp
# XXX FIXME TODO  Fix the GPG-to-SSH private key converter from github.com/pinpox/pgp2ssh
# XXX FIXME TODO  Generate QR codes for the newly generated keys and revcert(s)
# XXX FIXME TODO  Maybe this paperkey can help github.com/FiloSottile/mostly-harmless/tree/main/paper
# XXX FIXME TODO  Maybe combine this with stuff from github.com/tylert/file-fetcher/booya

# Can we use https://github.com/mikalv/anything2ed25519 to rebuild keys???
